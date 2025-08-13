package oauth

type User struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	Username     string `json:"username"`
	Email        string `json:"email"`
	RegisterDate int    `json:"registerDate"`
}

type UserSession struct {
	ID        string `json:"id"`
	ExpiresAt int    `json:"expiresAt"`
	UserId    int    `json:"userId"`
}
