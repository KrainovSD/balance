package oauthPlugin

import (
	"balance/internal/lib/api"
	"balance/internal/lib/web"
	"bytes"
	"database/sql"
	"errors"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

func callbackHandle(oauth *Oauth, options *OauthRegisterOptions) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var err error
		var status int = 500
		var proto = web.GetProto(r, oauth.frontendProtocol)
		var code = r.URL.Query().Get("code")
		var state = r.URL.Query().Get("state")
		var comebackUrlCookie *http.Cookie
		var comebackQuery url.Values
		var comebackUrl *url.URL
		var callbackUrl *http.Cookie
		var timeKey *http.Cookie
		var originState string

		var response api.Response
		var accessToken Token
		var formData url.Values
		var user User

		var sessionToken string
		var ownerId int
		var oauthId string
		var userId int

		/** Clear service cookies */
		http.SetCookie(w, &http.Cookie{
			Name:     oauth.cookieNameComebackUrl,
			Value:    "",
			Path:     options.AuthPath,
			Expires:  time.Unix(0, 0),
			HttpOnly: true,
			Secure:   proto == "https",
		})
		http.SetCookie(w, &http.Cookie{
			Name:     oauth.cookieNameTimeKey,
			Value:    "",
			Path:     options.AuthPath,
			Expires:  time.Unix(0, 0),
			HttpOnly: true,
			Secure:   proto == "https",
		})
		http.SetCookie(w, &http.Cookie{
			Name:     oauth.cookieNameCallbackUrl,
			Value:    "",
			Path:     options.AuthPath,
			Expires:  time.Unix(0, 0),
			HttpOnly: true,
			Secure:   proto == "https",
		})

		/** Extract and check service data */
		if comebackUrlCookie, err = r.Cookie(oauth.cookieNameComebackUrl); err != nil {
			goto FATAL
		}
		if comebackUrl, err = url.Parse(comebackUrlCookie.Value); err != nil {
			goto FATAL
		}
		comebackQuery = comebackUrl.Query()

		if callbackUrl, err = r.Cookie(oauth.cookieNameCallbackUrl); err != nil {
			goto FATAL
		}

		if timeKey, err = r.Cookie(oauth.cookieNameTimeKey); err != nil {
			goto FATAL
		}

		originState, err = oauth.redis.Get(timeKey.Value)
		if err != nil {
			goto FATAL
		}
		if originState != state {
			err = errors.New("the state is not the same")
			goto FATAL
		}

		/** Get Token */
		formData = url.Values{}
		formData.Set("grant_type", "authorization_code")
		formData.Set("client_id", options.ClientId)
		formData.Set("client_secret", options.ClientSecret)
		formData.Set("code", code)
		formData.Set("redirect_uri", callbackUrl.Value)

		if response, err = oauth.apiClient.Send(api.Request{
			Url:         options.TokenUrl,
			Method:      api.Methods.POST,
			ContentType: api.ContentTypes.Form,
			Body:        bytes.NewBufferString(formData.Encode()),
		}); err != nil {
			goto FATAL
		}

		if accessToken, err = options.ParseToken(response.Data); err != nil {
			goto FATAL
		}

		/** Get User */
		if response, err = oauth.apiClient.Send(api.Request{
			Url:         options.UserUrl,
			Method:      api.Methods.GET,
			ContentType: api.ContentTypes.JSON,
			Headers:     map[string]string{"Authorization": "Bearer " + accessToken.AccessToken},
		}); err != nil {
			goto FATAL
		}

		if user, err = options.ParseUser(response.Data); err != nil {
			goto FATAL
		}

		/** Check User In DB */

		oauthId = options.Provider + ":" + user.ID

		if ownerId, err = GetUserId(r); err == nil {
			if userId, err = oauth.usersProvider.GetUserIdByProvider(oauthId); err == nil && userId == ownerId {
				goto SUCCESS
			}

			if err = oauth.usersProvider.CreateProvider(ownerId, oauthId); err != nil {
				goto FATAL
			}

			userId = ownerId
			goto SUCCESS
		}

		if userId, err = oauth.usersProvider.GetUserIdByProvider(oauthId); err != nil {
			if err == sql.ErrNoRows {
				if userId, err = oauth.usersProvider.CreateUser(user.Name, user.UserName, user.Email, oauthId); err != nil {
					goto FATAL
				}
			} else {
				goto FATAL
			}
		}

		/** Generate and Set session token */
	SUCCESS:
		if sessionToken, err = oauth.usersProvider.CreateSession(userId, oauth.stateLength, oauth.sessionTokenExpires); err != nil {
			goto FATAL
		}
		oauth.usersProvider.UpdateLastLogin(userId)

		http.SetCookie(w, &http.Cookie{
			Name:     oauth.cookieNameToken,
			Value:    sessionToken,
			Path:     "/",
			Expires:  time.Now().Add(oauth.sessionTokenExpires),
			HttpOnly: true,
			Secure:   proto == "https",
		})

		comebackQuery.Set(oauth.queryExpires, strconv.FormatInt(time.Now().Add(oauth.sessionTokenExpires-time.Second*30).UnixMilli(), 10))
		comebackUrl.RawQuery = comebackQuery.Encode()
		http.Redirect(w, r, comebackUrl.String(), http.StatusTemporaryRedirect)
		return

	FATAL:
		web.SendError(w, web.ErrorResponse{
			Status: status,
			Error:  err,
		})
	}
}
