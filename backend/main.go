package main

import (
	"database/sql"
	"encoding/json"
	"finances/db"
	"finances/lib"
	"finances/oauth"
	"finances/plugins"
	"fmt"
	"net/http"
	"os"
	"time"
)

var m http.ServeMux

func main() {
	var err error
	var redis *plugins.RedisClient
	var apiClient *lib.ApiClient
	var pg *sql.DB

	if err = lib.LoadEnvFile(".env"); err != nil {
		panic(err.Error())
	}
	var cookieNameToken = os.Getenv("AUTH_COOKIE")
	if redis, err = plugins.CreateRedisClient(plugins.RedisClientOptions{
		Timeout: 5 * time.Second,
	}); err != nil {
		panic(err.Error())
	}
	if apiClient, err = lib.CreateApiClient(); err != nil {
		panic(err.Error())
	}
	if pg, err = plugins.CreatePgClient(); err != nil {
		panic(err.Error())
	}
	defer pg.Close()

	if err = db.InitDb(pg); err != nil {
		panic(err.Error())

	}

	if err = oauth.InitGitlabOauth(oauth.Oauth{
		M:                     &m,
		Redis:                 redis,
		Db:                    pg,
		ApiClient:             apiClient,
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
		panic(err.Error())
	}

	if err = oauth.InitGoogleOauth(oauth.Oauth{
		M:                     &m,
		Redis:                 redis,
		Db:                    pg,
		ApiClient:             apiClient,
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
		panic(err.Error())
	}

	testHandle := func(w http.ResponseWriter, r *http.Request) {
		var userID int
		var err error

		if userID, err = oauth.GetUserId(r); err != nil {
			lib.SendError(w, lib.ErrorResponse{
				Error: err,
			})
			return
		}

		type ExampleResponse struct {
			UserId int `json:"user_id"`
		}
		end := ExampleResponse{
			UserId: userID,
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(end)
	}

	var port = os.Getenv("PORT")
	m.Handle("/test", oauth.AuthMiddleware(pg, cookieNameToken, false)(http.HandlerFunc(testHandle)))

	fmt.Println("Starting Server on " + port)
	panic(http.ListenAndServe(":"+port, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
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
