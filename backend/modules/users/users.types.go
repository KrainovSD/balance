package users

type User struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	Username     string `json:"username"`
	Email        string `json:"email"`
	LastLogin    string `json:"lastLogin"`
	RegisterDate string `json:"registerDate"`
}

type UserSession struct {
	ID        string `json:"id"`
	ExpiresAt int    `json:"expiresAt"`
	UserId    int    `json:"userId"`
}
