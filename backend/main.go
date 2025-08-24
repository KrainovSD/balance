package main

import (
	"balance/internal/config"
	"balance/internal/db"
	"balance/internal/lib/api"
	"balance/internal/modules/users"
	oauthPlugin "balance/internal/plugins/oauth"
	pgPlugin "balance/internal/plugins/postgres"
	redisPlugin "balance/internal/plugins/redis"
	"balance/internal/routes"
	"database/sql"
	"fmt"
	"net/http"
	"time"
)

var m http.ServeMux

func main() {
	var err error
	var redis *redisPlugin.Client
	var apiClient *api.Client
	var pg *sql.DB
	var router *routes.Router
	var conf config.Config
	var auth *oauthPlugin.Auth
	var oauth *oauthPlugin.Oauth

	if conf, err = config.Create(); err != nil {
		panic(err.Error())
	}

	if redis, err = redisPlugin.CreateClient(redisPlugin.Options{
		Timeout:  5 * time.Second,
		Host:     conf.REDIS_HOST,
		Port:     conf.REDIS_PORT,
		Password: conf.REDIS_PASSWORD,
		Username: conf.REDIS_USERNAME,
	}); err != nil {
		panic(err.Error())
	}

	if apiClient, err = api.CreateClient(); err != nil {
		panic(err.Error())
	}
	if pg, err = pgPlugin.CreateClient(pgPlugin.Options{
		Username: conf.POSTGRES_USERNAME,
		Password: conf.POSTGRES_PASSWORD,
		Host:     conf.POSTGRES_HOST,
		Port:     conf.POSTGRES_PORT,
		Name:     conf.POSTGRES_DB,
		Ssl:      false,
	}); err != nil {
		panic(err.Error())
	}
	defer pg.Close()

	if err = db.InitTables(pg); err != nil {
		panic(err.Error())

	}

	if auth, err = oauthPlugin.CreateAuth(oauthPlugin.AuthOptions{
		UsersProvider: &users.UsersProvider{
			Db: pg,
		},
		CookieName: conf.AUTH_COOKIE,
	}); err != nil {
		panic(err.Error())
	}

	if oauth, err = oauthPlugin.Create(oauthPlugin.OauthOptions{
		M:         &m,
		Redis:     redis,
		ApiClient: apiClient,
		Db:        pg,
		UsersProvider: &users.UsersProvider{
			Db: pg,
		},
		CookieNameCallbackUrl: "balance.callback",
		CookieNameComebackUrl: "balance.comeback",
		CookieNameTimeKey:     "balance.timeKey",
		QueryExpires:          "session_token_expires",
		CookieNameToken:       conf.AUTH_COOKIE,
		StateLength:           16,
		ServiceDataExpires:    time.Minute * 5,
		SessionTokenExpires:   time.Hour * 24 * 5,
		FrontendHost:          conf.FRONTEND_HOST,
		FrontendProtocol:      conf.FRONTEND_PROTOCOL,
	}); err != nil {
		panic(err.Error())
	}

	router = routes.CreateRouter(routes.RouterOptions{
		Mux:       &m,
		Redis:     redis,
		Db:        pg,
		ApiClient: apiClient,
		UsersProvider: &users.UsersProvider{
			Db: pg,
		},
		Auth:   auth,
		Oauth:  oauth,
		Config: conf,
	})
	if err = router.Init(); err != nil {
		panic(err.Error())
	}

	fmt.Println("Starting Server on " + conf.PORT)
	panic(http.ListenAndServe(":"+conf.PORT, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("access-control-allow-origin", "*")
		w.Header().Set("access-control-allow-headers", "*")
		w.Header().Set("access-control-allow-methods", "*")
		w.Header().Set("access-control-allow", "*")
		if r.Method == "OPTIONS" {
			w.WriteHeader(200)
			return
		}
		m.ServeHTTP(w, r)
	})))

}
