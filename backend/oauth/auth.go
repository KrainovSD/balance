package oauth

import (
	"context"
	"database/sql"
	"errors"
	"finances/lib"
	"net/http"
	"strings"
	"time"
)

type contextKey string

var UserIDContextKey contextKey = "userID"

func AuthMiddleware(db *sql.DB, cookieName string, strict bool) func(http.Handler) http.Handler {
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
				if cookie, err = r.Cookie(cookieName); err != nil {
					if strict {
						goto FATAL
					} else {
						goto EXIT
					}
				}
				token = cookie.Value
			} else {
				token = strings.Replace(token, "Bearer ", "", 1)
			}

			if err = db.QueryRow("SELECT user_id, expires_at FROM balance.user_sessions WHERE id = $1", token).Scan(&userId, &expiresAt); err != nil {
				if err == sql.ErrNoRows {
					if strict {
						goto NOT_AUTHORIZED
					} else {
						goto EXIT

					}
				}
				if strict {
					goto FATAL
				} else {
					goto EXIT

				}
			}
			if time.Now().After(expiresAt) {
				if _, err = db.Exec("DELETE FROM balance.user_sessions WHERE id = $1", token); err != nil {
					if strict {
						goto FATAL
					} else {
						goto EXIT

					}
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
			lib.SendError(w, lib.ErrorResponse{
				Error:  err,
				Status: 401,
			})
			return

		FATAL:
			lib.SendError(w, lib.ErrorResponse{
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
