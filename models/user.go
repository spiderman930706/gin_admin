package models

type User struct {
	Model
	Username string `json:"username" list:"username" type:"string"`
	Password string `json:"password" type:"password"`
	IsAdmin bool `json:"is_admin" list:"is_admin" type:"bool"`
	IsStaff bool `json:"is_staff" list:"is_staff" type:"bool"`
}

