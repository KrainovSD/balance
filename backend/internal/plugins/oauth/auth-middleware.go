package oauthPlugin

import (
	"balance/internal/lib/web"
	"context"
	"database/sql"
	"errors"
	"net/http"
	"strings"
	"time"
)

type Auth struct {
	cookieName    string
	usersProvider userProvider
}

type AuthOptions struct {
	CookieName    string
	UsersProvider userProvider
}

func (a *AuthOptions) validate() error {
	if a.CookieName == "" {
		return errors.New("cookieName is empty")
	}

	return nil
}

func CreateAuth(options AuthOptions) (*Auth, error) {
	var err error

	if err = options.validate(); err != nil {
		return &Auth{}, err
	}

	return &Auth{
		cookieName:    options.CookieName,
		usersProvider: options.UsersProvider,
	}, nil
}

type contextKey string

var UserIDContextKey contextKey = "userID"

func (auth *Auth) Middleware(strict bool) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			var token string
			var cookie *http.Cookie
			var err error
			var userId int
			var expiresAt time.Time
			var ctx context.Context

			token = r.Header.Get("Authorization")
			if token == "" {
				if cookie, err = r.Cookie(auth.cookieName); err != nil {

					if strict {
						goto NOT_AUTHORIZED
					} else {
						goto EXIT
					}
				}
				token = cookie.Value
			} else {
				token = strings.Replace(token, "Bearer ", "", 1)
			}

			if userId, expiresAt, err = auth.usersProvider.GetSessionByToken(token); err != nil {
				if err == sql.ErrNoRows {
					if strict {
						goto NOT_AUTHORIZED
					} else {
						goto EXIT
					}
				}

				goto FATAL
			}
			if time.Now().After(expiresAt) {
				if err = auth.usersProvider.DeleteSession(token); err != nil {
					goto FATAL
				}

				if strict {
					goto NOT_AUTHORIZED
				} else {
					goto EXIT

				}
			}

			ctx = context.WithValue(r.Context(), UserIDContextKey, userId)
			next.ServeHTTP(w, r.WithContext(ctx))

			return

		EXIT:
			next.ServeHTTP(w, r)
			return

		NOT_AUTHORIZED:
			web.SendError(w, web.ErrorResponse{
				Error:  err,
				Status: 401,
			})
			return

		FATAL:
			web.SendError(w, web.ErrorResponse{
				Error: err,
			})

		})
	}
}

func GetUserId(r *http.Request) (int, error) {
	userID := r.Context().Value(UserIDContextKey)
	if userID == nil {
		return 0, errors.New("empty userID")
	}

	return userID.(int), nil

}
