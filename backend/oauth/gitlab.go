package oauth

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"finances/lib"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

type GitlabTokenResponse struct {
	AccessToken  string `json:"access_token"`
	TokenType    string `json:"token_type"`
	ExpiresIn    int    `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
}
type GitlabUserResponse struct {
	ID       int    `json:"id"`
	UserName string `json:"username"`
	Name     string `json:"name"`
	Email    string `json:"email"`
}

func InitGitlabOauth(oauth Oauth) error {
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
		authPath = "/api/v1/oauth/gitlab"
	}
	var callbackPath string
	if oauth.CallbackPath != "" {
		callbackPath = oauth.CallbackPath
	} else {
		callbackPath = "/api/v1/oauth/gitlab/callback"
	}

	oauth.M.HandleFunc(authPath, func(w http.ResponseWriter, r *http.Request) {
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
	})

	oauth.M.HandleFunc(callbackPath, func(w http.ResponseWriter, r *http.Request) {
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
		var response lib.Response
		var accessToken GitlabTokenResponse
		formData := url.Values{}
		formData.Set("grant_type", "authorization_code")
		formData.Set("client_id", env.ClientId)
		formData.Set("client_secret", env.ClientSecret)
		formData.Set("code", callbackServiceData.Code)
		formData.Set("redirect_uri", callbackServiceData.CallbackUrl)

		if response, err = oauth.ApiClient.Send(lib.Request{
			Url:         env.TokenUrl,
			Method:      lib.Methods.POST,
			ContentType: lib.ContentTypes.Form,
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
		var user GitlabUserResponse

		if response, err = oauth.ApiClient.Send(lib.Request{
			Url:         env.UserUrl,
			Method:      lib.Methods.GET,
			ContentType: lib.ContentTypes.JSON,
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
		var userId int
		oauthId := "gitlab:" + strconv.Itoa(user.ID)
		if err = oauth.Db.QueryRow(`SELECT id from finances.users WHERE oauth_id = $1`, oauthId).Scan(&userId); err != nil {
			if err == sql.ErrNoRows {
				if err = oauth.Db.QueryRow(`INSERT INTO finances.users (oauth_id, name, username, email) VALUES ($1, $2, $3, $4) returning id`, oauthId, user.Name, user.UserName, user.Email).Scan(&userId); err != nil {
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
		if sessionToken, err = lib.RandomHex(16); err != nil {
			lib.SendError(w, lib.ErrorResponse{
				Message: "generate session token",
				Error:   err,
			})
			return
		}
		if _, err = oauth.Db.Exec(`INSERT INTO finances.user_sessions (id, user_id, expires_at) VALUES ($1, $2, $3)`, sessionToken, userId, time.Now().Add(oauth.SessionTokenExpires)); err != nil {
			lib.SendError(w, lib.ErrorResponse{
				Message: "set session token to db. INSERT INTO finances.user_sessions (id, user_id, expires_at) VALUES " + sessionToken + " " + strconv.Itoa(userId),
				Error:   err,
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
	})

	return nil
}
