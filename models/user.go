package models

type User struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
	email    string `json:"email"`
	password string `json:"password"`
}
