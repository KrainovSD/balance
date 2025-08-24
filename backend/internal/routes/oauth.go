package routes

import (
	"balance/internal/config"
	oauthPlugin "balance/internal/plugins/oauth"
	"encoding/json"
	"strconv"
)

type oauthPreset struct {
	Config config.Config
}

func (o *oauthPreset) CreateGithub() oauthPlugin.OauthRegisterOptions {
	return oauthPlugin.OauthRegisterOptions{
		AuthPath:     "/api/v1/oauth/github",
		CallbackPath: "/api/v1/oauth/github/callback",
		ClientId:     o.Config.GITHUB_OAUTH_CLIENT_ID,
		ClientSecret: o.Config.GITHUB_OAUTH_CLIENT_SECRET,
		LoginUrl:     o.Config.GITHUB_OAUTH_LOGIN_URL,
		TokenUrl:     o.Config.GITHUB_OAUTH_TOKEN_URL,
		UserUrl:      o.Config.GITHUB_OAUTH_USER_URL,
		Provider:     "github",
		ParseUser: func(response []byte) (oauthPlugin.User, error) {
			var user oauthPlugin.GithubUser
			var err error
			if err = json.Unmarshal(response, &user); err != nil {

				return oauthPlugin.User{}, err
			}
			return oauthPlugin.User{ID: strconv.Itoa(user.ID), UserName: user.UserName, Name: user.UserName, Email: user.Name}, nil
		},
		ParseToken: func(response []byte) (oauthPlugin.Token, error) {
			var token oauthPlugin.GithubToken
			var err error
			if err = json.Unmarshal(response, &token); err != nil {

				return oauthPlugin.Token{}, err
			}
			return oauthPlugin.Token{AccessToken: token.AccessToken}, nil
		},
		Scopes: []string{"user"},
	}
}

func (o *oauthPreset) CreateGitlab() oauthPlugin.OauthRegisterOptions {
	return oauthPlugin.OauthRegisterOptions{
		AuthPath:     "/api/v1/oauth/gitlab",
		CallbackPath: "/api/v1/oauth/gitlab/callback",
		ClientId:     o.Config.GITLAB_OAUTH_CLIENT_ID,
		ClientSecret: o.Config.GITLAB_OAUTH_CLIENT_SECRET,
		LoginUrl:     o.Config.GITLAB_OAUTH_LOGIN_URL,
		TokenUrl:     o.Config.GITLAB_OAUTH_TOKEN_URL,
		UserUrl:      o.Config.GITLAB_OAUTH_USER_URL,
		Provider:     "gitlab",
		ParseUser: func(response []byte) (oauthPlugin.User, error) {
			var user oauthPlugin.GitlabUser
			var err error
			if err = json.Unmarshal(response, &user); err != nil {

				return oauthPlugin.User{}, err
			}
			return oauthPlugin.User{ID: strconv.Itoa(user.ID), UserName: user.UserName, Name: user.UserName, Email: user.Name}, nil
		},
		ParseToken: func(response []byte) (oauthPlugin.Token, error) {
			var token oauthPlugin.GitlabToken
			var err error
			if err = json.Unmarshal(response, &token); err != nil {

				return oauthPlugin.Token{}, err
			}
			return oauthPlugin.Token{AccessToken: token.AccessToken}, nil
		},
		Scopes: []string{"openid", "profile", "read_user", "email"},
	}
}

func (o *oauthPreset) CreateGoogle() oauthPlugin.OauthRegisterOptions {
	return oauthPlugin.OauthRegisterOptions{
		AuthPath:     "/api/v1/oauth/google",
		CallbackPath: "/api/v1/oauth/google/callback",
		ClientId:     o.Config.GOOGLE_OAUTH_CLIENT_ID,
		ClientSecret: o.Config.GOOGLE_OAUTH_CLIENT_SECRET,
		LoginUrl:     o.Config.GOOGLE_OAUTH_LOGIN_URL,
		TokenUrl:     o.Config.GOOGLE_OAUTH_TOKEN_URL,
		UserUrl:      o.Config.GOOGLE_OAUTH_USER_URL,
		Provider:     "google",
		ParseUser: func(response []byte) (oauthPlugin.User, error) {
			var user oauthPlugin.GoogleUser
			var err error
			if err = json.Unmarshal(response, &user); err != nil {

				return oauthPlugin.User{}, err
			}
			return oauthPlugin.User{ID: user.ID, UserName: user.UserName, Name: user.UserName, Email: user.Name}, nil
		},
		ParseToken: func(response []byte) (oauthPlugin.Token, error) {
			var token oauthPlugin.GoogleToken
			var err error
			if err = json.Unmarshal(response, &token); err != nil {

				return oauthPlugin.Token{}, err
			}
			return oauthPlugin.Token{AccessToken: token.AccessToken}, nil
		},
		Scopes: []string{"openid", "profile", "email"},
	}
}

func (o *oauthPreset) CreateYandex() oauthPlugin.OauthRegisterOptions {
	return oauthPlugin.OauthRegisterOptions{
		AuthPath:     "/api/v1/oauth/yandex",
		CallbackPath: "/api/v1/oauth/yandex/callback",
		ClientId:     o.Config.YANDEX_OAUTH_CLIENT_ID,
		ClientSecret: o.Config.YANDEX_OAUTH_CLIENT_SECRET,
		LoginUrl:     o.Config.YANDEX_OAUTH_LOGIN_URL,
		TokenUrl:     o.Config.YANDEX_OAUTH_TOKEN_URL,
		UserUrl:      o.Config.YANDEX_OAUTH_USER_URL,
		Provider:     "yandex",
		ParseUser: func(response []byte) (oauthPlugin.User, error) {
			var user oauthPlugin.YandexUser
			var err error
			if err = json.Unmarshal(response, &user); err != nil {

				return oauthPlugin.User{}, err
			}
			return oauthPlugin.User{ID: user.ID, UserName: user.UserName, Name: user.UserName, Email: user.Name}, nil
		},
		ParseToken: func(response []byte) (oauthPlugin.Token, error) {
			var token oauthPlugin.YandexToken
			var err error
			if err = json.Unmarshal(response, &token); err != nil {

				return oauthPlugin.Token{}, err
			}
			return oauthPlugin.Token{AccessToken: token.AccessToken}, nil
		},
		Scopes: []string{"login:email", "login:info"},
	}
}
