package db

import (
	"database/sql"
)

func InitDb(db *sql.DB) error {
	var err error

	if _, err = db.Exec(`
	CREATE TABLE IF NOT EXISTS finances.users (
		id SERIAL PRIMARY KEY,
		oauth_id VARCHAR NOT NULL,
		name VARCHAR NOT NULL,
		username VARCHAR NOT NULL,
		email VARCHAR,
		register_date TIMESTAMPTZ DEFAULT NOW()
	)
		`); err != nil {
		return err
	}

	if _, err = db.Exec(`
	CREATE TABLE IF NOT EXISTS finances.user_sessions (
		id VARCHAR PRIMARY KEY,
		user_id INT REFERENCES finances.users(id),
		expires_at TIMESTAMPTZ NOT NULL
	)
	`); err != nil {
		return err
	}

	return nil

}
