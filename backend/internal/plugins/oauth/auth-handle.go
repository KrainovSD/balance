package oauthPlugin

import (
	"balance/internal/lib/helpers"
	"balance/internal/lib/web"
	"net/http"
	"net/url"
	"strings"
	"time"
)

func authHandle(oauth *Oauth, options *OauthRegisterOptions) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var err error
		var status int = 500
		var proto = web.GetProto(r, oauth.frontendProtocol)
		var host = web.GetHost(r, oauth.frontendHost)
		var state string
		var timeKey string
		var comebackUrl = r.URL.Query().Get("comebackUrl")
		var callbackUrl = proto + "://" + host + options.CallbackPath
		var loginUrl *url.URL
		var query url.Values

		comebackUrl = strings.Replace(comebackUrl, proto+"://", "", 1)
		comebackUrl = strings.Replace(comebackUrl, host, "", 1)
		comebackUrl = proto + "://" + host + comebackUrl

		if state, err = helpers.RandomHex(oauth.stateLength); err != nil {
			goto FATAL
		}
		if timeKey, err = helpers.RandomHex(oauth.stateLength); err != nil {
			goto FATAL
		}
		if err = oauth.redis.Set(timeKey, state, oauth.serviceDataExpires); err != nil {
			goto FATAL
		}

		/** Set service cookies */
		http.SetCookie(w, &http.Cookie{
			Name:     oauth.cookieNameComebackUrl,
			Value:    comebackUrl,
			Path:     options.AuthPath,
			Expires:  time.Now().Add(oauth.serviceDataExpires),
			HttpOnly: true,
			Secure:   proto == "https",
		})
		http.SetCookie(w, &http.Cookie{
			Name:     oauth.cookieNameTimeKey,
			Value:    timeKey,
			Path:     options.AuthPath,
			Expires:  time.Now().Add(oauth.serviceDataExpires),
			HttpOnly: true,
			Secure:   proto == "https",
		})
		http.SetCookie(w, &http.Cookie{
			Name:     oauth.cookieNameCallbackUrl,
			Value:    callbackUrl,
			Path:     options.AuthPath,
			Expires:  time.Now().Add(oauth.serviceDataExpires),
			HttpOnly: true,
			Secure:   proto == "https",
		})

		/** Generate oauth url */
		if loginUrl, err = url.Parse(options.LoginUrl); err != nil {
			goto FATAL
		}
		query = loginUrl.Query()
		query.Add("state", state)
		query.Add("client_id", options.ClientId)
		query.Add("response_type", "code")
		query.Add("redirect_uri", callbackUrl)
		if len(options.Scopes) > 0 {
			query.Add("scope", strings.Join(options.Scopes, " "))
		}
		loginUrl.RawQuery = query.Encode()

		http.Redirect(w, r, loginUrl.String(), http.StatusTemporaryRedirect)
		return

	FATAL:
		web.SendError(w, web.ErrorResponse{
			Status: status,
			Error:  err,
		})

	}

}
