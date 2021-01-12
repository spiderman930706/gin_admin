package models

type User struct {
	Model
	Username string `gorm:"type:varchar(15);unique_index" json:"username" admin:"list;name:用户名"`
	Password string `json:"password" admin:"type:password;name:密码"`
	IsAdmin  bool   `json:"is_admin" admin:"list;name:超级管理员"`
	IsStaff  bool   `json:"is_staff" admin:"list;name:职员"`
}
