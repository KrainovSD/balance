package routes

import (
	"database/sql"
	"finances/api"
	"finances/oauth"
	"finances/plugins"
	"net/http"
	"os"
	"time"
)

type Auth struct {
	Mux       *http.ServeMux
	Redis     *plugins.RedisClient
	Db        *sql.DB
	ApiClient *api.ApiClient
}
type AuthRouter interface {
	Init() error
}

func CreateAuthRouter(auth Auth) AuthRouter {
	return &auth
}

func (a *Auth) Init() error {
	var err error
	var cookieNameToken = os.Getenv("AUTH_COOKIE")

	if err = oauth.InitGitlabOauth(oauth.Oauth{
		M:                     a.Mux,
		Redis:                 a.Redis,
		Db:                    a.Db,
		ApiClient:             a.ApiClient,
		AuthPath:              "/api/v1/oauth/gitlab",
		CallbackPath:          "/api/v1/oauth/gitlab/callback",
		ServiceDataExpires:    5 * time.Minute,
		SessionTokenExpires:   24 * time.Hour * 1,
		StateLength:           16,
		CookieNameCallbackUrl: "balance.callback",
		CookieNameComebackUrl: "balance.comeback",
		CookieNameTimeKey:     "balance.key",
		CookieNameToken:       cookieNameToken,
		PrefixEnv:             "GITLAB",
		Scopes:                []string{"openid", "profile", "read_user", "email"},
	}); err != nil {
		return err

	}

	if err = oauth.InitGoogleOauth(oauth.Oauth{
		M:                     a.Mux,
		Redis:                 a.Redis,
		Db:                    a.Db,
		ApiClient:             a.ApiClient,
		AuthPath:              "/api/v1/oauth/google",
		CallbackPath:          "/api/v1/oauth/google/callback",
		ServiceDataExpires:    5 * time.Minute,
		SessionTokenExpires:   24 * time.Hour * 1,
		StateLength:           16,
		CookieNameCallbackUrl: "balance.callback",
		CookieNameComebackUrl: "balance.comeback",
		CookieNameTimeKey:     "balance.key",
		CookieNameToken:       cookieNameToken,
		PrefixEnv:             "GOOGLE",
		Scopes:                []string{"openid", "profile", "email"},
	}); err != nil {
		return err

	}

	if err = oauth.InitYandexOauth(oauth.Oauth{
		M:                     a.Mux,
		Redis:                 a.Redis,
		Db:                    a.Db,
		ApiClient:             a.ApiClient,
		AuthPath:              "/api/v1/oauth/yandex",
		CallbackPath:          "/api/v1/oauth/yandex/callback",
		ServiceDataExpires:    5 * time.Minute,
		SessionTokenExpires:   24 * time.Hour * 1,
		StateLength:           16,
		CookieNameCallbackUrl: "balance.callback",
		CookieNameComebackUrl: "balance.comeback",
		CookieNameTimeKey:     "balance.key",
		CookieNameToken:       cookieNameToken,
		PrefixEnv:             "YANDEX",
		Scopes:                []string{"login:email", "login:info"},
	}); err != nil {
		return err
	}

	if err = oauth.InitGithubOauth(oauth.Oauth{
		M:                     a.Mux,
		Redis:                 a.Redis,
		Db:                    a.Db,
		ApiClient:             a.ApiClient,
		AuthPath:              "/api/v1/oauth/github",
		CallbackPath:          "/api/v1/oauth/github/callback",
		ServiceDataExpires:    5 * time.Minute,
		SessionTokenExpires:   24 * time.Hour * 1,
		StateLength:           16,
		CookieNameCallbackUrl: "balance.callback",
		CookieNameComebackUrl: "balance.comeback",
		CookieNameTimeKey:     "balance.key",
		CookieNameToken:       cookieNameToken,
		PrefixEnv:             "GITHUB",
		Scopes:                []string{"user"},
	}); err != nil {
		return err
	}

	return err
}
