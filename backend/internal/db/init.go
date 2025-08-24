package db

import (
	"database/sql"
)

func InitTables(db *sql.DB) error {
	var err error

	if err = initUsers(db); err != nil {
		return err
	}
	if err = initReceipts(db); err != nil {
		return err
	}
	if err = initPayments(db); err != nil {
		return err
	}

	return err
}

func initUsers(db *sql.DB) error {
	var err error

	if _, err = db.Exec(`
	CREATE TABLE IF NOT EXISTS balance.users (
		id SERIAL PRIMARY KEY,
		name VARCHAR NOT NULL,
		username VARCHAR NOT NULL,
		email VARCHAR,
		last_login TIMESTAMPTZ DEFAULT NOW(),
		register_date TIMESTAMPTZ DEFAULT NOW()
	)
		`); err != nil {
		return err
	}

	if _, err = db.Exec(`
	CREATE TABLE IF NOT EXISTS balance.user_sessions (
		id VARCHAR PRIMARY KEY,
		user_id INT REFERENCES balance.users(id) ON DELETE CASCADE ON UPDATE CASCADE,
		expires_at TIMESTAMPTZ NOT NULL
	)
	`); err != nil {
		return err
	}

	if _, err = db.Exec(`
	CREATE TABLE IF NOT EXISTS balance.user_providers (
		id VARCHAR PRIMARY KEY,
		user_id INT REFERENCES balance.users(id) ON DELETE CASCADE ON UPDATE CASCADE
	)
	`); err != nil {
		return err
	}

	return nil
}

func initReceipts(db *sql.DB) error {
	var err error

	if _, err = db.Exec(`
	CREATE TABLE IF NOT EXISTS balance.receipt_templates (
    	id SERIAL PRIMARY KEY,
    	name VARCHAR NOT NULL,
    	amount INT NOT NULL,
    	user_id INT REFERENCES balance.users(id) ON DELETE CASCADE ON UPDATE CASCADE
    );
	`); err != nil {
		return err
	}

	if _, err = db.Exec(`
	CREATE TABLE IF NOT EXISTS balance.receipts (
    	id SERIAL PRIMARY KEY,
    	amount INT NOT NULL,
    	description VARCHAR,
		receipt_id INT REFERENCES balance.receipt_templates(id),
    	user_id INT REFERENCES balance.users(id) ON DELETE CASCADE ON UPDATE CASCADE,
    	date TIMESTAMPTZ DEFAULT NOW()
		
);
	`); err != nil {
		return err
	}

	return err
}

func initPayments(db *sql.DB) error {
	var err error

	if _, err = db.Exec(`
	CREATE TABLE IF NOT EXISTS balance.payment_templates (
    	id SERIAL PRIMARY KEY,
    	name VARCHAR NOT NULL,
    	amount INT NOT NULL,
    	user_id INT REFERENCES balance.users(id) ON DELETE CASCADE ON UPDATE CASCADE
	);
	`); err != nil {
		return err
	}

	if _, err = db.Exec(`
	CREATE TABLE IF NOT EXISTS balance.payments (
    	id SERIAL PRIMARY KEY,
    	amount INT NOT NULL,
    	description VARCHAR,
    	payment_id INT REFERENCES balance.payment_templates(id),
    	user_id INT REFERENCES balance.users(id) ON DELETE CASCADE ON UPDATE CASCADE,
    	date TIMESTAMPTZ DEFAULT NOW()
	);
	`); err != nil {
		return err
	}

	return err
}
