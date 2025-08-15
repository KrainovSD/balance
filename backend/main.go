package main

import (
	"database/sql"
	"encoding/json"
	"finances/api"
	"finances/db"
	"finances/lib"
	"finances/oauth"
	"finances/plugins"
	"finances/routes"
	"fmt"
	"net/http"
	"os"
	"time"
)

var m http.ServeMux

func main() {
	var err error
	var redis *plugins.RedisClient
	var apiClient *api.ApiClient
	var pg *sql.DB

	if err = lib.LoadEnvFile(".env"); err != nil {
		panic(err.Error())
	}
	if redis, err = plugins.CreateRedisClient(plugins.RedisClientOptions{
		Timeout: 5 * time.Second,
	}); err != nil {
		panic(err.Error())
	}
	if apiClient, err = api.CreateApiClient(); err != nil {
		panic(err.Error())
	}
	if pg, err = plugins.CreatePgClient(); err != nil {
		panic(err.Error())
	}
	defer pg.Close()

	if err = db.InitDb(pg); err != nil {
		panic(err.Error())

	}

	authRouter := routes.CreateAuthRouter(routes.Auth{
		Mux:       &m,
		Redis:     redis,
		Db:        pg,
		ApiClient: apiClient,
	})
	if err = authRouter.Init(); err != nil {
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
	var cookieNameToken = os.Getenv("AUTH_COOKIE")
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
