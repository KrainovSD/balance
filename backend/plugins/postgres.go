package plugins

import (
	"database/sql"
	"errors"
	"os"
	"time"

	_ "github.com/lib/pq"
)

func CreatePgClient() (*sql.DB, error) {
	user := os.Getenv("POSTGRES_USERNAME")
	if user == "" {
		return nil, errors.New("hasn't POSTGRES_USERNAME")
	}
	password := os.Getenv("POSTGRES_PASSWORD")
	if password == "" {
		return nil, errors.New("hasn't POSTGRES_PASSWORD")
	}
	host := os.Getenv("POSTGRES_HOST")
	if host == "" {
		return nil, errors.New("hasn't POSTGRES_HOST")
	}
	port := os.Getenv("POSTGRES_PORT")
	if port == "" {
		return nil, errors.New("hasn't POSTGRES_PORT")
	}
	dbName := os.Getenv("POSTGRES_DB")
	if dbName == "" {
		return nil, errors.New("hasn't POSTGRES_DB")
	}

	var db *sql.DB
	var err error

	connect := "host=" + host + " port=" + port + " user=" + user + " password=" + password + " dbname=" + dbName + " sslmode=disable connect_timeout=10"
	if db, err = sql.Open("postgres", connect); err != nil {
		return db, err
	}

	if err = db.Ping(); err != nil {
		return db, err
	}

	db.SetConnMaxLifetime(5 * time.Minute)
	db.SetConnMaxIdleTime(5 * time.Minute)
	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(5)

	return db, err
}
