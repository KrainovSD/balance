package oauth

import (
	"finances/lib"
	"net/http"
	"net/url"
	"strings"
)

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
		query.Add("state", serviceInfo.TimeKey)
		query.Add("client_id", env.ClientId)
		query.Add("response_type", "code")
		query.Add("redirect_url", callbackUrl)
		if len(oauth.Scopes) > 0 {
			query.Add("scope", strings.Join(oauth.Scopes, " "))
		}
		loginUrl.RawQuery = query.Encode()

		http.Redirect(w, r, loginUrl.String(), http.StatusTemporaryRedirect)
	})

	oauth.M.HandleFunc(callbackPath, func(w http.ResponseWriter, r *http.Request) {

	})

	return nil
}
