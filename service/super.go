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

func ChangeRoleAuth(roleID int, batchID models.BatchID) (err error) {
	role := models.Role{
		Model: models.Model{ID: uint(roleID)},
	}
	var authList []models.Auth
	var auth models.Auth
	global.DB.Model(&auth).Where("id in ?", batchID.IDList).Find(&authList)
	err = global.DB.Model(&role).Association("Auth").Replace(authList)
	return
}

func GetAuthList() (authList []models.AuthList, err error) {
	var auth models.Auth
	global.DB.Model(&auth).Find(&authList)
	return
}
