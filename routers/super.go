package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/spiderman930706/gin_admin/api"
)

func InitSuperUserRouter(Router *gin.RouterGroup) {
	apiRouter := Router.Group("super")
	{
		apiRouter.GET("/role/:role_id", api.RoleAuthList)
	}
}
