package users

import (
	"balance/internal/lib/web"
	oauthPlugin "balance/internal/plugins/oauth"
	"database/sql"
	"encoding/json"
	"net/http"
)

type UserControllerOptions struct {
	Db   *sql.DB
	Auth *oauthPlugin.Auth
}
type UserInterface interface {
	Init(mux *http.ServeMux)
}

func CreateUserController(options UserControllerOptions) UserInterface {
	return &UsersController{
		UsersService: &UsersService{
			UsersProvider: &UsersProvider{
				Db: options.Db,
			},
		},
		Auth: options.Auth,
	}
}

type UsersController struct {
	UsersService *UsersService
	Auth         *oauthPlugin.Auth
}

func (r *UsersController) GetUser(w http.ResponseWriter, req *http.Request) {
	var user User
	var userID int
	var err error
	userID, _ = oauthPlugin.GetUserId(req)

	if user, err = r.UsersService.GetUser(userID); err != nil {
		web.SendError(w, web.ErrorResponse{
			Error: err,
		})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)

}

func (r *UsersController) Init(mux *http.ServeMux) {

	mux.Handle("/api/v1/users",
		r.Auth.Middleware(true)(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
			switch req.Method {

			case "GET":
				r.GetUser(w, req)

			default:
				w.WriteHeader(405)
			}
		})))

}
