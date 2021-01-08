package routers

import (
	"github.com/gin-gonic/gin"

	"gin_admin/api"
)

func InitUserRouter(Router *gin.RouterGroup) {
	apiRouter := Router.Group("user")
	{
		apiRouter.GET("", api.GetAdminTableList)
	}
}
