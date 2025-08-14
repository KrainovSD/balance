package db

import (
	"database/sql"
)

func InitDb(db *sql.DB) error {
	var err error

	if _, err = db.Exec(`
	CREATE TABLE IF NOT EXISTS balance.users (
		id SERIAL PRIMARY KEY,
		name VARCHAR NOT NULL,
		username VARCHAR NOT NULL,
		email VARCHAR,
		register_date TIMESTAMPTZ DEFAULT NOW()
	)
		`); err != nil {
		return err
	}

	if _, err = db.Exec(`
	CREATE TABLE IF NOT EXISTS balance.user_sessions (
		id VARCHAR PRIMARY KEY,
		user_id INT REFERENCES balance.users(id),
		expires_at TIMESTAMPTZ NOT NULL
	)
	`); err != nil {
		return err
	}

	if _, err = db.Exec(`
	CREATE TABLE IF NOT EXISTS balance.user_providers (
		id VARCHAR PRIMARY KEY,
		user_id INT REFERENCES balance.users(id)
	)
	`); err != nil {
		return err
	}

	return nil

}
