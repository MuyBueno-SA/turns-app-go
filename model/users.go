package model

type User struct {
	ID       int    `json:"id"`
	Username string `json:"name"`
	Phone    string `json:"phone"`
	Email    string `json:"email"`
	Activity string `json:"activity"`
}
