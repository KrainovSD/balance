package oauthPlugin

type GithubToken struct {
	AccessToken string `json:"access_token"`
}
type GitlabToken struct {
	AccessToken string `json:"access_token"`
}
type GoogleToken struct {
	AccessToken string `json:"access_token"`
}
type YandexToken struct {
	AccessToken string `json:"access_token"`
}

type GithubUser struct {
	ID       int    `json:"id"`
	UserName string `json:"login"`
	Name     string `json:"name"`
	Email    string `json:"email"`
}
type GitlabUser struct {
	ID       int    `json:"id"`
	UserName string `json:"username"`
	Name     string `json:"name"`
	Email    string `json:"email"`
}
type GoogleUser struct {
	ID       string `json:"id"`
	UserName string `json:"family_name"`
	Name     string `json:"name"`
	Email    string `json:"email"`
}
type YandexUser struct {
	ID       string `json:"id"`
	UserName string `json:"last_name"`
	Name     string `json:"real_name"`
	Email    string `json:"default_email"`
}
