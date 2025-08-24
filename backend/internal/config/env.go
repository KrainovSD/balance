package config

import (
	"balance/internal/lib/helpers"
	"errors"
	"os"
)

type Config struct {
	PORT              string
	FRONTEND_PROTOCOL string
	FRONTEND_HOST     string
	AUTH_COOKIE       string

	REDIS_HOST     string
	REDIS_PORT     string
	REDIS_PASSWORD string
	REDIS_USERNAME string

	POSTGRES_HOST     string
	POSTGRES_PORT     string
	POSTGRES_PASSWORD string
	POSTGRES_USERNAME string
	POSTGRES_DB       string

	GOOGLE_OAUTH_LOGIN_URL     string
	GOOGLE_OAUTH_TOKEN_URL     string
	GOOGLE_OAUTH_USER_URL      string
	GOOGLE_OAUTH_CLIENT_SECRET string
	GOOGLE_OAUTH_CLIENT_ID     string

	YANDEX_OAUTH_LOGIN_URL     string
	YANDEX_OAUTH_TOKEN_URL     string
	YANDEX_OAUTH_USER_URL      string
	YANDEX_OAUTH_CLIENT_SECRET string
	YANDEX_OAUTH_CLIENT_ID     string

	GITLAB_OAUTH_LOGIN_URL     string
	GITLAB_OAUTH_TOKEN_URL     string
	GITLAB_OAUTH_USER_URL      string
	GITLAB_OAUTH_CLIENT_SECRET string
	GITLAB_OAUTH_CLIENT_ID     string

	GITHUB_OAUTH_LOGIN_URL     string
	GITHUB_OAUTH_TOKEN_URL     string
	GITHUB_OAUTH_USER_URL      string
	GITHUB_OAUTH_CLIENT_SECRET string
	GITHUB_OAUTH_CLIENT_ID     string
}

func Create() (Config, error) {
	var config Config
	var err error

	if err = helpers.LoadEnvFile(".env"); err != nil {
		return config, err
	}

	config.FRONTEND_PROTOCOL = os.Getenv("FRONTEND_PROTOCOL")
	config.FRONTEND_HOST = os.Getenv("FRONTEND_HOST")
	config.REDIS_PASSWORD = os.Getenv("REDIS_PASSWORD")
	config.REDIS_USERNAME = os.Getenv("REDIS_USERNAME")
	config.POSTGRES_PASSWORD = os.Getenv("POSTGRES_PASSWORD")
	config.POSTGRES_USERNAME = os.Getenv("POSTGRES_USERNAME")
	config.POSTGRES_DB = os.Getenv("POSTGRES_DB")
	config.GOOGLE_OAUTH_LOGIN_URL = os.Getenv("GOOGLE_OAUTH_LOGIN_URL")
	config.GOOGLE_OAUTH_TOKEN_URL = os.Getenv("GOOGLE_OAUTH_TOKEN_URL")
	config.GOOGLE_OAUTH_USER_URL = os.Getenv("GOOGLE_OAUTH_USER_URL")
	config.GOOGLE_OAUTH_CLIENT_SECRET = os.Getenv("GOOGLE_OAUTH_CLIENT_SECRET")
	config.GOOGLE_OAUTH_CLIENT_ID = os.Getenv("GOOGLE_OAUTH_CLIENT_ID")
	config.YANDEX_OAUTH_LOGIN_URL = os.Getenv("YANDEX_OAUTH_LOGIN_URL")
	config.YANDEX_OAUTH_TOKEN_URL = os.Getenv("YANDEX_OAUTH_TOKEN_URL")
	config.YANDEX_OAUTH_USER_URL = os.Getenv("YANDEX_OAUTH_USER_URL")
	config.YANDEX_OAUTH_CLIENT_SECRET = os.Getenv("YANDEX_OAUTH_CLIENT_SECRET")
	config.YANDEX_OAUTH_CLIENT_ID = os.Getenv("YANDEX_OAUTH_CLIENT_ID")
	config.GITLAB_OAUTH_LOGIN_URL = os.Getenv("GITLAB_OAUTH_LOGIN_URL")
	config.GITLAB_OAUTH_TOKEN_URL = os.Getenv("GITLAB_OAUTH_TOKEN_URL")
	config.GITLAB_OAUTH_USER_URL = os.Getenv("GITLAB_OAUTH_USER_URL")
	config.GITLAB_OAUTH_CLIENT_SECRET = os.Getenv("GITLAB_OAUTH_CLIENT_SECRET")
	config.GITLAB_OAUTH_CLIENT_ID = os.Getenv("GITLAB_OAUTH_CLIENT_ID")
	config.GITHUB_OAUTH_LOGIN_URL = os.Getenv("GITHUB_OAUTH_LOGIN_URL")
	config.GITHUB_OAUTH_TOKEN_URL = os.Getenv("GITHUB_OAUTH_TOKEN_URL")
	config.GITHUB_OAUTH_USER_URL = os.Getenv("GITHUB_OAUTH_USER_URL")
	config.GITHUB_OAUTH_CLIENT_SECRET = os.Getenv("GITHUB_OAUTH_CLIENT_SECRET")
	config.GITHUB_OAUTH_CLIENT_ID = os.Getenv("GITHUB_OAUTH_CLIENT_ID")

	if os.Getenv("PORT") == "" {
		return config, errors.New("hasn't env PORT")
	}
	config.PORT = os.Getenv("PORT")

	if os.Getenv("AUTH_COOKIE") == "" {
		return config, errors.New("hasn't env AUTH_COOKIE")
	}
	config.AUTH_COOKIE = os.Getenv("AUTH_COOKIE")

	if os.Getenv("REDIS_HOST") == "" {
		return config, errors.New("hasn't env REDIS_HOST")
	}
	config.REDIS_HOST = os.Getenv("REDIS_HOST")

	if os.Getenv("REDIS_PORT") == "" {
		return config, errors.New("hasn't env REDIS_PORT")
	}
	config.REDIS_PORT = os.Getenv("REDIS_PORT")

	if os.Getenv("POSTGRES_HOST") == "" {
		return config, errors.New("hasn't env POSTGRES_HOST")
	}
	config.POSTGRES_HOST = os.Getenv("POSTGRES_HOST")

	if os.Getenv("POSTGRES_PORT") == "" {
		return config, errors.New("hasn't env POSTGRES_PORT")
	}
	config.POSTGRES_PORT = os.Getenv("POSTGRES_PORT")

	return config, nil
}
