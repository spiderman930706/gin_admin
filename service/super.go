package service

import (
	"github.com/spiderman930706/gin_admin/global"
	"github.com/spiderman930706/gin_admin/models"
)

func GetRoleAuthList(roleID int) (role models.Role, err error) {
	result := global.DB.Where("id = ?", roleID).Preload("Auth").First(&role)
	err = result.Error
	return
}
