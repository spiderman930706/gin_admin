package models

type User struct {
	Model
	Username string `gorm:"type:varchar(15);uniqueIndex;not null" json:"username" admin:"list;name:用户名"`
	Password string `json:"password" admin:"type:password;name:密码" gorm:"not null"`
	IsAdmin  bool   `json:"is_admin" admin:"list;name:超级管理员" gorm:"default:false;not null"`
	IsStaff  bool   `json:"is_staff" admin:"list;name:职员" gorm:"default:false;not null"`
	Role     Role
	RoleID   int `json:"role_id" gorm:"index" admin:"name:角色"`
}

type Role struct {
	Model
	Name string `gorm:"unique;not null" admin:"name:角色名称;list"`
	Auth []Auth `gorm:"many2many:role_auths;"`
}

type Auth struct {
	Model
	Table  string `json:"table" gorm:"uniqueIndex:table_method;not null"`
	Method string `json:"method" gorm:"uniqueIndex:table_method;not null"`
}
