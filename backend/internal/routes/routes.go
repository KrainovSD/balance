package routes

import (
	"balance/internal/config"
	"balance/internal/lib/api"
	"balance/internal/modules/payments"
	"balance/internal/modules/receipts"
	"balance/internal/modules/users"
	oauthPlugin "balance/internal/plugins/oauth"
	"database/sql"
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
	CreateUser(name string, username string, email string, provider string) (int, error)
	GetUserIdByProvider(provider string) (int, error)
	UpdateLastLogin(userID int) error
}

type RouterOptions struct {
	Mux           *http.ServeMux
	Redis         redisClient
	Db            *sql.DB
	ApiClient     *api.Client
	UsersProvider userProvider
	Auth          *oauthPlugin.Auth
	Oauth         *oauthPlugin.Oauth
	Config        config.Config
}

type Router struct {
	auth          *oauthPlugin.Auth
	oauth         *oauthPlugin.Oauth
	mux           *http.ServeMux
	redis         redisClient
	db            *sql.DB
	apiClient     *api.Client
	usersProvider userProvider
	config        config.Config
}

func CreateRouter(options RouterOptions) *Router {
	return &Router{
		mux:           options.Mux,
		redis:         options.Redis,
		db:            options.Db,
		apiClient:     options.ApiClient,
		auth:          options.Auth,
		oauth:         options.Oauth,
		usersProvider: options.UsersProvider,
		config:        options.Config,
	}
}

func (r *Router) Init() error {
	var err error
	var oauthPreset oauthPreset = oauthPreset{Config: r.config}

	r.oauth.Register(oauthPreset.CreateGithub())
	r.oauth.Register(oauthPreset.CreateGitlab())
	r.oauth.Register(oauthPreset.CreateGoogle())
	r.oauth.Register(oauthPreset.CreateYandex())

	paymentController := payments.CreatePaymentController(payments.PaymentControllerOptions{
		Db:   r.db,
		Auth: r.auth,
	})
	paymentController.Init(r.mux)

	receiptController := receipts.CreateReceiptController(receipts.ReceiptControllerOptions{
		Db:   r.db,
		Auth: r.auth,
	})
	receiptController.Init(r.mux)

	userController := users.CreateUserController(users.UserControllerOptions{
		Db:   r.db,
		Auth: r.auth,
	})
	userController.Init(r.mux)

	return err

}
