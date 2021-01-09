package service

import (
	"errors"

	"github.com/spiderman930706/gin_admin"
	"github.com/spiderman930706/gin_admin/models"
)

func GetUserId(id int) (err error, user *models.User) {
	var u models.User
	if err = gin_admin.DB.Where("`id` = ?", id).First(&u).Error; err != nil {
		return errors.New("用户不存在"), &u
	}
	return nil, &u
}