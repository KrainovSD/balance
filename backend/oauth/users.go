package oauth

import (
	"database/sql"
	"finances/lib"
	"time"
)

type UsersProvider struct {
	Db *sql.DB
}

type IUsersProvider interface {
	GetUserIdByProvider(provider string) (int, error)
	CreateProvider(userID int, provider string) error
	CreateUser(user User, provider string) (int, error)
	CreateSession(userID int, length int, expires time.Duration) (string, error)
}

func (u *UsersProvider) GetUserIdByProvider(provider string) (int, error) {
	var userID int
	err := u.Db.QueryRow(`SELECT user_id from balance.user_providers WHERE id = $1`, provider).Scan(&userID)

	return userID, err
}

func (u *UsersProvider) CreateProvider(userID int, provider string) error {
	_, err := u.Db.Exec(`INSERT INTO balance.user_providers (id, user_id) VALUES ($1, $2)`, provider, userID)

	return err
}

func (u *UsersProvider) CreateUser(user User, provider string) (int, error) {
	var userID int
	var err error
	var tx *sql.Tx

	if tx, err = u.Db.Begin(); err != nil {
		return userID, err
	}

	defer func() {
		if err != nil {
			tx.Rollback()
			return
		}
		err = tx.Commit()
	}()

	if err = tx.QueryRow(`INSERT INTO balance.users (name, username, email) VALUES ($1, $2, $3) returning id`, user.Name, user.Username, user.Email).Scan(&userID); err != nil {
		return userID, err
	}
	_, err = tx.Exec(`INSERT INTO balance.user_providers (id, user_id) VALUES ($1, $2)`, provider, userID)

	return userID, err
}

func (u *UsersProvider) CreateSession(userID int, length int, expires time.Duration) (string, error) {
	var sessionToken string
	var err error

	if sessionToken, err = lib.RandomHex(length); err != nil {
		return sessionToken, err
	}
	_, err = u.Db.Exec(`INSERT INTO balance.user_sessions (id, user_id, expires_at) VALUES ($1, $2, $3)`, sessionToken, userID, time.Now().Add(expires))

	return sessionToken, err
}
