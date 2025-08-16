package oauth

import (
	"database/sql"
	"errors"
	"finances/api"
	"finances/lib"
	"net/http"
	"os"
	"strings"
	"time"
)

type RedisClient interface {
	Get(key string) (string, error)
	Set(key string, value string, expiration time.Duration) error
	Del(key ...string) error
}

type Oauth struct {
	M                     *http.ServeMux
	Redis                 RedisClient
	ApiClient             *api.ApiClient
	Db                    *sql.DB
	AuthPath              string
	CallbackPath          string
	ServiceDataExpires    time.Duration
	SessionTokenExpires   time.Duration
	StateLength           int
	CookieNameCallbackUrl string
	CookieNameComebackUrl string
	CookieNameTimeKey     string
	CookieNameToken       string
	PrefixEnv             string
	Scopes                []string
}

func (o *Oauth) Validate() error {
	if o.AuthPath == "" {
		return errors.New("hasn't auth path in oauth with prefix" + o.PrefixEnv)
	}
	if o.CallbackPath == "" {
		return errors.New("hasn't callback path in oauth with prefix" + o.PrefixEnv)
	}
	if o.ServiceDataExpires == 0 {
		o.ServiceDataExpires = 5 * time.Minute
	}
	if o.StateLength == 0 {
		o.StateLength = 16
	}
	if o.CookieNameCallbackUrl == "" {
		return errors.New("hasn't cookie name callback in oauth with prefix" + o.PrefixEnv)
	}
	if o.CookieNameComebackUrl == "" {
		return errors.New("hasn't cookie name comeback in oauth with prefix" + o.PrefixEnv)
	}
	if o.CookieNameTimeKey == "" {
		return errors.New("hasn't cookie name time key in oauth with prefix" + o.PrefixEnv)
	}

	return nil
}

type OauthServiceData struct {
	Proto       string
	Host        string
	Secure      bool
	State       string
	TimeKey     string
	ComebackUrl string
}

func (o *Oauth) GetServiceInfo(r *http.Request) (OauthServiceData, error) {
	var serviceInfo OauthServiceData
	var err error
	var proto = lib.GetProto(r)
	var host = lib.GetHost(r)
	var secure bool
	if proto == "https" {
		secure = true
	}
	var comebackUrl = r.URL.Query().Get("comebackUrl")
	comebackUrl = strings.Replace(comebackUrl, proto+"://", "", 1)
	comebackUrl = strings.Replace(comebackUrl, host, "", 1)
	var state string
	var timeKey string

	if state, err = lib.RandomHex(o.StateLength); err != nil {
		return serviceInfo, err
	}
	if timeKey, err = lib.RandomHex(o.StateLength); err != nil {
		return serviceInfo, err
	}

	serviceInfo = OauthServiceData{
		Proto:       proto,
		Host:        host,
		Secure:      secure,
		State:       state,
		TimeKey:     timeKey,
		ComebackUrl: proto + "://" + host + comebackUrl,
	}

	return serviceInfo, nil
}

type OauthEnv struct {
	ClientId     string
	ClientSecret string
	LoginUrl     string
	TokenUrl     string
	UserUrl      string
}

func (o *Oauth) GetEnv() (OauthEnv, error) {
	var err error
	var env OauthEnv

	var clientId = os.Getenv(o.PrefixEnv + "_OAUTH_CLIENT_ID")
	var clientSecret = os.Getenv(o.PrefixEnv + "_OAUTH_CLIENT_SECRET")
	var loginUrl = os.Getenv(o.PrefixEnv + "_OAUTH_LOGIN_URL")
	var tokenUrl = os.Getenv(o.PrefixEnv + "_OAUTH_TOKEN_URL")
	var userUrl = os.Getenv(o.PrefixEnv + "_OAUTH_USER_URL")

	if clientId == "" {
		err = errors.New("couldn't get _OAUTH_CLIENT_ID with prefix" + o.PrefixEnv)
		return env, err
	}
	if clientSecret == "" {
		err = errors.New("couldn't get _OAUTH_CLIENT_SECRET with prefix" + o.PrefixEnv)
		return env, err
	}
	if loginUrl == "" {
		err = errors.New("couldn't get _OAUTH_LOGIN_URL with prefix" + o.PrefixEnv)
		return env, err
	}
	if tokenUrl == "" {
		err = errors.New("couldn't get _OAUTH_TOKEN_URL with prefix" + o.PrefixEnv)
		return env, err
	}
	if userUrl == "" {
		err = errors.New("couldn't get _OAUTH_USER_URL with prefix" + o.PrefixEnv)
		return env, err
	}

	env = OauthEnv{
		ClientId:     clientId,
		ClientSecret: clientSecret,
		LoginUrl:     loginUrl,
		TokenUrl:     tokenUrl,
		UserUrl:      userUrl,
	}

	return env, nil
}

type OauthCookie struct {
	ComebackUrl string
	CallbackUrl string
	TimeKey     string
	Clear       bool
	Secure      bool
}

func (o *Oauth) SetServiceCookie(w http.ResponseWriter, c OauthCookie) {
	var expires time.Time
	if c.Clear {
		expires = time.Unix(0, 0)
	} else {
		expires = time.Now().Add(o.ServiceDataExpires)
	}

	http.SetCookie(w, &http.Cookie{
		Name:     o.CookieNameComebackUrl,
		Value:    c.ComebackUrl,
		Path:     o.AuthPath,
		Expires:  expires,
		HttpOnly: true,
		Secure:   c.Secure,
	})
	http.SetCookie(w, &http.Cookie{
		Name:     o.CookieNameTimeKey,
		Value:    c.TimeKey,
		Path:     o.AuthPath,
		Expires:  expires,
		HttpOnly: true,
		Secure:   c.Secure,
	})
	http.SetCookie(w, &http.Cookie{
		Name:     o.CookieNameCallbackUrl,
		Value:    c.CallbackUrl,
		Path:     o.AuthPath,
		Expires:  expires,
		HttpOnly: true,
		Secure:   c.Secure,
	})
}

type CallbackServiceData struct {
	ComebackUrl string
	CallbackUrl string
	Code        string
}

func (o *Oauth) GetCallbackServiceData(r *http.Request) (CallbackServiceData, error) {
	var data CallbackServiceData
	var err error
	var code = r.URL.Query().Get("code")
	var state = r.URL.Query().Get("state")
	var comebackUrl *http.Cookie
	var callbackUrl *http.Cookie
	var timeKey *http.Cookie
	var originState string

	if comebackUrl, err = r.Cookie(o.CookieNameComebackUrl); err != nil {
		return data, err
	}

	if callbackUrl, err = r.Cookie(o.CookieNameCallbackUrl); err != nil {
		return data, err
	}

	if timeKey, err = r.Cookie(o.CookieNameTimeKey); err != nil {
		return data, err
	}

	originState, err = o.Redis.Get(timeKey.Value)
	if err != nil {
		return data, err
	}
	if originState != state {
		err = errors.New("the state is not the same")
		return data, err
	}

	data = CallbackServiceData{
		ComebackUrl: comebackUrl.Value,
		CallbackUrl: callbackUrl.Value,
		Code:        code,
	}

	return data, nil
}
