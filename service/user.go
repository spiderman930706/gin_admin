package service

import (
	"github.com/spiderman930706/gin_admin/global"
	"github.com/spiderman930706/gin_admin/models"
	"github.com/spiderman930706/gin_admin/util"
)

func Login(u *models.User) (err error, userInter *models.User) {
	var user models.User
	u.Password = util.MD5V([]byte(u.Password))
	err = global.DB.Where("username = ? AND password = ?", u.Username, u.Password).Preload("Role").First(&user).Error
	return err, &user
}
