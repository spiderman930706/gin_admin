package api

import (
	"github.com/gin-gonic/gin"
	"github.com/spiderman930706/gin_admin/models"
	"github.com/spiderman930706/gin_admin/service"
	"strconv"
)

func RoleAuthList(c *gin.Context) {
	roleIdStr := c.Param("role_id")
	roleID, err := strconv.Atoi(roleIdStr)
	if err != nil {
		FailWithMessage(err.Error(), c)
		return
	}
	if roleData, err := service.GetRoleAuthList(roleID); err != nil {
		FailWithMessage("获取失败", c)
	} else {
		OkWithDetailed(roleData, "获取成功", c)
	}
}

func ChangeRoleAuth(c *gin.Context) {
	roleIdStr := c.Param("role_id")
	roleID, err := strconv.Atoi(roleIdStr)
	if err != nil {
		FailWithMessage(err.Error(), c)
		return
	}
	var idList models.BatchID
	if err := c.BindJSON(&idList); err != nil {
		FailWithMessage("请求参数有误", c)
		return
	}
	if err := service.ChangeRoleAuth(roleID, idList); err != nil {
		FailWithMessage("修改失败", c)
	} else {
		OkWithMessage("修改成功", c)
	}
}

func AuthList(c *gin.Context) {
	if authList, err := service.GetAuthList(); err != nil {
		FailWithMessage("修改失败", c)
	} else {
		OkWithData(authList, c)
	}
}
