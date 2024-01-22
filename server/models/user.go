package models

type User struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	Password string `json:"password"`
	Admin    bool   `json:"admin"`
	ApiKey   string `json:"apikey"`
}
