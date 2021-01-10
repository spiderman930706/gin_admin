package models

type User struct {
	Model
	Username string `gorm:"type:varchar(15);unique_index" json:"username" admin:"list:username"`
	Password string `json:"password" admin:"list:username;type:password"`
	IsAdmin  bool   `json:"is_admin" admin:"list:is_admin"`
	IsStaff  bool   `json:"is_staff" admin:"list:is_staff"`
}
