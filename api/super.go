package api

import (
	"github.com/gin-gonic/gin"
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
