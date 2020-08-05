package model

type User struct {
	ID       float64 `json:"id"`
	Username string  `json:"username"`
	Name     string  `json:"name"`
	Phone    string  `json:"phone"`
	Email    string  `json:"email"`
	Avatar   string  `json:"avatar"`
}
