package routes

import (
	"database/sql"
	"finances/api"
	"finances/modules/payments"
	"finances/modules/receipts"
	"finances/oauth"
	"finances/plugins"
	"net/http"
	"time"
)

type Router struct {
	Mux             *http.ServeMux
	Redis           *plugins.RedisClient
	Db              *sql.DB
	ApiClient       *api.ApiClient
	cookieNameToken string
}

func (r *Router) Init() error {
	var err error

	if err = oauth.InitGitlabOauth(oauth.Oauth{
		M:                     r.Mux,
		Redis:                 r.Redis,
		Db:                    r.Db,
		ApiClient:             r.ApiClient,
		AuthPath:              "/api/v1/oauth/gitlab",
		CallbackPath:          "/api/v1/oauth/gitlab/callback",
		ServiceDataExpires:    5 * time.Minute,
		SessionTokenExpires:   24 * time.Hour * 1,
		StateLength:           16,
		CookieNameCallbackUrl: "balance.callback",
		CookieNameComebackUrl: "balance.comeback",
		CookieNameTimeKey:     "balance.key",
		CookieNameToken:       r.cookieNameToken,
		PrefixEnv:             "GITLAB",
		Scopes:                []string{"openid", "profile", "read_user", "email"},
	}); err != nil {
		return err

	}

	if err = oauth.InitGoogleOauth(oauth.Oauth{
		M:                     r.Mux,
		Redis:                 r.Redis,
		Db:                    r.Db,
		ApiClient:             r.ApiClient,
		AuthPath:              "/api/v1/oauth/google",
		CallbackPath:          "/api/v1/oauth/google/callback",
		ServiceDataExpires:    5 * time.Minute,
		SessionTokenExpires:   24 * time.Hour * 1,
		StateLength:           16,
		CookieNameCallbackUrl: "balance.callback",
		CookieNameComebackUrl: "balance.comeback",
		CookieNameTimeKey:     "balance.key",
		CookieNameToken:       r.cookieNameToken,
		PrefixEnv:             "GOOGLE",
		Scopes:                []string{"openid", "profile", "email"},
	}); err != nil {
		return err

	}

	if err = oauth.InitYandexOauth(oauth.Oauth{
		M:                     r.Mux,
		Redis:                 r.Redis,
		Db:                    r.Db,
		ApiClient:             r.ApiClient,
		AuthPath:              "/api/v1/oauth/yandex",
		CallbackPath:          "/api/v1/oauth/yandex/callback",
		ServiceDataExpires:    5 * time.Minute,
		SessionTokenExpires:   24 * time.Hour * 1,
		StateLength:           16,
		CookieNameCallbackUrl: "balance.callback",
		CookieNameComebackUrl: "balance.comeback",
		CookieNameTimeKey:     "balance.key",
		CookieNameToken:       r.cookieNameToken,
		PrefixEnv:             "YANDEX",
		Scopes:                []string{"login:email", "login:info"},
	}); err != nil {
		return err
	}

	if err = oauth.InitGithubOauth(oauth.Oauth{
		M:                     r.Mux,
		Redis:                 r.Redis,
		Db:                    r.Db,
		ApiClient:             r.ApiClient,
		AuthPath:              "/api/v1/oauth/github",
		CallbackPath:          "/api/v1/oauth/github/callback",
		ServiceDataExpires:    5 * time.Minute,
		SessionTokenExpires:   24 * time.Hour * 1,
		StateLength:           16,
		CookieNameCallbackUrl: "balance.callback",
		CookieNameComebackUrl: "balance.comeback",
		CookieNameTimeKey:     "balance.key",
		CookieNameToken:       r.cookieNameToken,
		PrefixEnv:             "GITHUB",
		Scopes:                []string{"user"},
	}); err != nil {
		return err
	}

	receipts := receipts.ReceiptController{
		ReceiptService: receipts.ReceiptService{
			ReceiptProvider: receipts.ReceiptProvider{
				Db: r.Db,
			},
		},
		CookieNameToken: r.cookieNameToken,
	}
	receipts.Init(r.Mux)
	payments := payments.PaymentController{
		PaymentService: payments.PaymentService{
			PaymentProvider: payments.PaymentProvider{
				Db: r.Db,
			},
		},
		CookieNameToken: r.cookieNameToken,
	}
	payments.Init(r.Mux)

	return err
}
