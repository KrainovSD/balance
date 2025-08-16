package users

import (
	"encoding/json"
	"finances/lib"
	"finances/oauth"
	"net/http"
)

type UsersController struct {
	UsersService    UsersService
	CookieNameToken string
}

func (r *UsersController) GetUser(w http.ResponseWriter, req *http.Request) {
	var user User
	var userID int
	var err error
	userID, _ = oauth.GetUserId(req)

	if user, err = r.UsersService.GetUser(userID); err != nil {
		lib.SendError(w, lib.ErrorResponse{
			Error: err,
		})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)

}

func (r *UsersController) Init(mux *http.ServeMux) {

	mux.Handle("/api/v1/users", oauth.AuthMiddleware(r.UsersService.UsersProvider.Db, r.CookieNameToken, true)(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		switch req.Method {

		case "GET":
			r.GetUser(w, req)

		default:
			w.WriteHeader(405)
		}
	})))

}
