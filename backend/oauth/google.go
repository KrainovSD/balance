package oauth

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"finances/api"
	"finances/lib"
	"net/http"
	"net/url"
	"strings"
	"time"
)

type GoogleTokenResponse struct {
	AccessToken string `json:"access_token"`
}

type GoogleUserResponse struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	Email      string `json:"email"`
	GivenName  string `json:"given_name"`
	FamilyName string `json:"family_name"`
}

func InitGoogleOauth(oauth Oauth) error {
	var env OauthEnv
	var err error

	if err = oauth.Validate(); err != nil {
		return err
	}

	if env, err = oauth.GetEnv(); err != nil {
		return err
	}

	var authPath string
	if oauth.AuthPath != "" {
		authPath = oauth.AuthPath
	} else {
		authPath = "/api/v1/oauth/google"
	}
	var callbackPath string
	if oauth.CallbackPath != "" {
		callbackPath = oauth.CallbackPath
	} else {
		callbackPath = "/api/v1/oauth/google/callback"
	}

	usersProvider := UsersProvider{
		Db: oauth.Db,
	}

	authHandle := func(w http.ResponseWriter, r *http.Request) {
		var err error
		var serviceInfo OauthServiceData

		if serviceInfo, err = oauth.GetServiceInfo(r); err != nil {
			lib.SendError(w, lib.ErrorResponse{
				Message: "generate service info",
				Error:   err,
			})
			return
		}
		if err = oauth.Redis.Set(serviceInfo.TimeKey, serviceInfo.State, oauth.ServiceDataExpires); err != nil {
			lib.SendError(w, lib.ErrorResponse{
				Message: "set state to redis",
				Error:   err,
			})
			return
		}
		callbackUrl := serviceInfo.Proto + "://" + serviceInfo.Host + callbackPath

		oauth.SetServiceCookie(w, OauthCookie{
			ComebackUrl: serviceInfo.ComebackUrl,
			CallbackUrl: callbackUrl,
			TimeKey:     serviceInfo.TimeKey,
			Clear:       false,
			Secure:      serviceInfo.Proto == "https",
		})

		var loginUrl *url.URL
		if loginUrl, err = url.Parse(env.LoginUrl); err != nil {
			lib.SendError(w, lib.ErrorResponse{
				Message: "parse login url",
				Error:   err,
			})
			return
		}
		query := loginUrl.Query()
		query.Add("state", serviceInfo.State)
		query.Add("client_id", env.ClientId)
		query.Add("response_type", "code")
		query.Add("redirect_uri", callbackUrl)
		if len(oauth.Scopes) > 0 {
			query.Add("scope", strings.Join(oauth.Scopes, " "))
		}
		loginUrl.RawQuery = query.Encode()

		http.Redirect(w, r, loginUrl.String(), http.StatusTemporaryRedirect)
	}
	callbackHandle := func(usersProvider IUsersProvider) func(w http.ResponseWriter, r *http.Request) {
		return func(w http.ResponseWriter, r *http.Request) {
			var proto = lib.GetProto(r)
			var callbackServiceData CallbackServiceData
			var err error
			oauth.SetServiceCookie(w, OauthCookie{
				Clear:  true,
				Secure: proto == "https",
			})

			if callbackServiceData, err = oauth.GetCallbackServiceData(r); err != nil {
				lib.SendError(w, lib.ErrorResponse{
					Message: "get service data",
					Error:   err,
				})
				return
			}

			/** Get Token */
			var response api.Response
			var accessToken GoogleTokenResponse
			formData := url.Values{}
			formData.Set("grant_type", "authorization_code")
			formData.Set("client_id", env.ClientId)
			formData.Set("client_secret", env.ClientSecret)
			formData.Set("code", callbackServiceData.Code)
			formData.Set("redirect_uri", callbackServiceData.CallbackUrl)

			if response, err = oauth.ApiClient.Send(api.Request{
				Url:         env.TokenUrl,
				Method:      api.Methods.POST,
				ContentType: api.ContentTypes.Form,
				Body:        bytes.NewBufferString(formData.Encode()),
			}); err != nil {
				lib.SendError(w, lib.ErrorResponse{
					Message: "get token",
					Error:   err,
				})
				return
			}

			if err = json.Unmarshal(response.Data, &accessToken); err != nil {
				lib.SendError(w, lib.ErrorResponse{
					Message: "parse token",
					Error:   err,
				})
				return
			}

			/** Get User */
			var user GoogleUserResponse

			if response, err = oauth.ApiClient.Send(api.Request{
				Url:         env.UserUrl,
				Method:      api.Methods.GET,
				ContentType: api.ContentTypes.JSON,
				Headers:     map[string]string{"Authorization": "Bearer " + accessToken.AccessToken},
			}); err != nil {
				lib.SendError(w, lib.ErrorResponse{
					Message: "get user",
					Error:   err,
				})
				return
			}

			if err = json.Unmarshal(response.Data, &user); err != nil {
				lib.SendError(w, lib.ErrorResponse{
					Message: "parse user",
					Error:   err,
				})
				return
			}

			/** Check User In DB */
			var ownerId int
			var userId int
			oauthId := "google:" + user.ID

			if ownerId, err = GetUserId(r); err == nil {
				if _, err = usersProvider.GetUserIdByProvider(oauthId); err == nil {
					http.Redirect(w, r, callbackServiceData.ComebackUrl, http.StatusTemporaryRedirect)
					return
				}

				if err = usersProvider.CreateProvider(ownerId, oauthId); err != nil {
					lib.SendError(w, lib.ErrorResponse{
						Message: "add new provider",
						Error:   err,
					})
					return
				}
				http.Redirect(w, r, callbackServiceData.ComebackUrl, http.StatusTemporaryRedirect)
				return
			}

			if userId, err = usersProvider.GetUserIdByProvider(oauthId); err != nil {
				if err == sql.ErrNoRows {
					if userId, err = usersProvider.CreateUser(User{
						Name:     user.Name,
						Username: user.FamilyName,
						Email:    user.Email,
					}, oauthId); err != nil {
						lib.SendError(w, lib.ErrorResponse{
							Message: "create user",
							Error:   err,
						})
						return
					}
				} else {
					lib.SendError(w, lib.ErrorResponse{
						Message: "get user from db",
						Error:   err,
					})
					return
				}
			}

			/** Generate and Set session token */
			var sessionToken string
			if sessionToken, err = usersProvider.CreateSession(userId, oauth.StateLength, oauth.SessionTokenExpires); err != nil {
				lib.SendError(w, lib.ErrorResponse{
					Error: err,
				})
				return
			}

			http.SetCookie(w, &http.Cookie{
				Name:     oauth.CookieNameToken,
				Value:    sessionToken,
				Path:     "/",
				Expires:  time.Now().Add(oauth.SessionTokenExpires),
				HttpOnly: true,
				Secure:   proto == "https",
			})
			http.Redirect(w, r, callbackServiceData.ComebackUrl, http.StatusTemporaryRedirect)
		}
	}

	oauth.M.HandleFunc(authPath, authHandle)
	oauth.M.Handle(callbackPath, AuthMiddleware(oauth.Db, oauth.CookieNameToken, false)(http.HandlerFunc(callbackHandle(&usersProvider))))

	return nil
}
