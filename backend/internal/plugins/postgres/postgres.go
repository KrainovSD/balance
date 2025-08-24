package pgPlugin

import (
	"database/sql"
	"errors"
	"fmt"
	"time"

	_ "github.com/lib/pq"
)

type Options struct {
	Username        string
	Password        string
	Host            string
	Port            string
	Name            string
	Ssl             bool
	ConnectTimeout  int
	ConnMaxLifetime time.Duration
	ConnMaxIdleTime time.Duration
	MaxIdleConns    int
	MaxOpenConns    int
}

func CreateClient(options Options) (*sql.DB, error) {
	if options.Host == "" {
		return nil, errors.New("hasn't POSTGRES_HOST")
	}
	if options.Port == "" {
		return nil, errors.New("hasn't POSTGRES_PORT")
	}
	connect := fmt.Sprintf("host=%s port=%s", options.Host, options.Port)
	if options.Username != "" {
		connect += fmt.Sprintf(" user=%s", options.Username)
	}
	if options.Password != "" {
		connect += fmt.Sprintf(" password=%s", options.Password)
	}
	if options.Name != "" {
		connect += fmt.Sprintf(" dbname=%s", options.Name)
	}

	if options.Ssl {
		connect += " sslmode=enable"
	} else {
		connect += " sslmode=disable"
	}

	if options.ConnectTimeout != 0 {
		connect += fmt.Sprintf(" connect_timeout=%d", options.ConnectTimeout)
	} else {
		connect += " connect_timeout=10"
	}

	var db *sql.DB
	var err error

	if db, err = sql.Open("postgres", connect); err != nil {
		return db, err
	}

	if err = db.Ping(); err != nil {
		return db, err
	}

	if options.ConnMaxLifetime != 0 {
		db.SetConnMaxLifetime(options.ConnMaxLifetime)
	} else {
		db.SetConnMaxLifetime(5 * time.Minute)
	}

	if options.ConnMaxIdleTime != 0 {
		db.SetConnMaxIdleTime(options.ConnMaxIdleTime)
	} else {
		db.SetConnMaxIdleTime(5 * time.Minute)
	}

	if options.MaxIdleConns != 0 {
		db.SetMaxIdleConns(options.MaxIdleConns)
	} else {
		db.SetMaxIdleConns(5)
	}

	if options.MaxOpenConns != 0 {
		db.SetMaxOpenConns(options.MaxOpenConns)
	} else {
		db.SetMaxOpenConns(5)
	}

	return db, err
}
