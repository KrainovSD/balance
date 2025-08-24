package oauthPlugin

import (
	"balance/internal/lib/api"
	"database/sql"
	"errors"
	"net/http"
	"time"
)

type redisClient interface {
	Get(key string) (string, error)
	Set(key string, value string, expiration time.Duration) error
	Del(key ...string) error
}

type userProvider interface {
	CreateProvider(userID int, provider string) error
	CreateSession(userID int, length int, expires time.Duration) (string, error)
	DeleteSession(token string) error
	GetSessionByToken(token string) (int, time.Time, error)
	CreateUser(name string, username string, email string, provider string) (int, error)
	GetUserIdByProvider(provider string) (int, error)
	UpdateLastLogin(userID int) error
}

type Oauth struct {
	m                     *http.ServeMux
	redis                 redisClient
	apiClient             *api.Client
	db                    *sql.DB
	usersProvider         userProvider
	auth                  *Auth
	cookieNameCallbackUrl string
	cookieNameComebackUrl string
	cookieNameTimeKey     string
	cookieNameToken       string
	queryExpires          string
	stateLength           int
	serviceDataExpires    time.Duration
	sessionTokenExpires   time.Duration
	frontendHost          string
	frontendProtocol      string
}

type OauthOptions struct {
	M                     *http.ServeMux
	Redis                 redisClient
	ApiClient             *api.Client
	Db                    *sql.DB
	UsersProvider         userProvider
	CookieNameCallbackUrl string
	CookieNameComebackUrl string
	CookieNameTimeKey     string
	CookieNameToken       string
	QueryExpires          string
	StateLength           int
	ServiceDataExpires    time.Duration
	SessionTokenExpires   time.Duration
	FrontendHost          string
	FrontendProtocol      string
}

func (o *OauthOptions) validate() error {
	if o == nil {
		return errors.New("oauthOptions pointer is nil")
	}
	if o.CookieNameCallbackUrl == "" {
		return errors.New("cookieNameCallbackUrl is empty")
	}
	if o.CookieNameComebackUrl == "" {
		return errors.New("cookieNameComebackUrl is empty")
	}
	if o.CookieNameTimeKey == "" {
		return errors.New("cookieNameTimeKey is empty")
	}
	if o.CookieNameToken == "" {
		return errors.New("cookieNameToken is empty")
	}
	if o.StateLength == 0 {
		return errors.New("stateLength is empty")
	}
	if o.ServiceDataExpires == 0 {
		return errors.New("serviceDataExpires is empty")
	}
	if o.SessionTokenExpires == 0 {
		return errors.New("sessionTokenExpires is empty")
	}

	return nil
}

func Create(options OauthOptions) (*Oauth, error) {
	var err error

	if err = options.validate(); err != nil {
		return nil, err
	}

	return &Oauth{
		m:         options.M,
		redis:     options.Redis,
		apiClient: options.ApiClient,
		db:        options.Db,
		auth: &Auth{
			cookieName:    options.CookieNameToken,
			usersProvider: options.UsersProvider,
		},
		usersProvider:         options.UsersProvider,
		cookieNameCallbackUrl: options.CookieNameCallbackUrl,
		cookieNameComebackUrl: options.CookieNameComebackUrl,
		cookieNameTimeKey:     options.CookieNameTimeKey,
		cookieNameToken:       options.CookieNameToken,
		queryExpires:          options.QueryExpires,
		stateLength:           options.StateLength,
		serviceDataExpires:    options.ServiceDataExpires,
		sessionTokenExpires:   options.SessionTokenExpires,
		frontendHost:          options.FrontendHost,
		frontendProtocol:      options.FrontendProtocol,
	}, nil
}

type Token struct {
	AccessToken string `json:"access_token"`
}
type User struct {
	ID       string `json:"id"`
	UserName string `json:"username"`
	Name     string `json:"name"`
	Email    string `json:"email"`
}

type OauthRegisterOptions struct {
	AuthPath     string
	CallbackPath string
	ClientId     string
	ClientSecret string
	LoginUrl     string
	TokenUrl     string
	UserUrl      string
	Provider     string
	ParseUser    func(response []byte) (User, error)
	ParseToken   func(response []byte) (Token, error)
	Scopes       []string
}

func (o *OauthRegisterOptions) validate() error {
	if o == nil {
		return errors.New("oauthRegisterOptions pointer is nil")
	}
	if o.AuthPath == "" {
		return errors.New("authPath is empty")
	}
	if o.CallbackPath == "" {
		return errors.New("callbackPath is empty")
	}
	if o.ClientId == "" {
		return errors.New("clientId is empty")
	}
	if o.ClientSecret == "" {
		return errors.New("clientSecret is empty")
	}
	if o.LoginUrl == "" {
		return errors.New("loginUrl is empty")
	}
	if o.TokenUrl == "" {
		return errors.New("tokenUrl is empty")
	}
	if o.UserUrl == "" {
		return errors.New("userUrl is empty")
	}
	if o.Provider == "" {
		return errors.New("provider is empty")
	}

	return nil
}

func (o *Oauth) Register(options OauthRegisterOptions) error {
	var err error

	if o == nil {
		return errors.New("Oauth pointer is nil")
	}
	if err = options.validate(); err != nil {
		return err
	}

	o.m.HandleFunc(options.AuthPath, authHandle(o, &options))
	o.m.Handle(options.CallbackPath, o.auth.Middleware(false)(http.HandlerFunc(callbackHandle(o, &options))))

	return nil
}
