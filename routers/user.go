package routers

import (
	"github.com/gin-gonic/gin"

	"github.com/spiderman930706/gin_admin/api"
)

func InitUserRouter(Router *gin.RouterGroup) {
	apiRouter := Router.Group("user")
	{
		apiRouter.GET("", api.GetAdminTableList)
	}
}
