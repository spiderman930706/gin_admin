package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/spiderman930706/gin_admin/api"
)

func InitPubUserRouter(Router *gin.RouterGroup) {
	apiRouter := Router.Group("user")
	{
		apiRouter.POST("login", api.UserLogin)
	}
}

//func InitAuthUserRouter(Router *gin.RouterGroup) {
//	apiRouter := Router.Group("auth_user")
//	{
//		apiRouter.POST("login", api.UserLogin)
//	}
//}
