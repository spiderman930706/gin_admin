package routers

import "github.com/gin-gonic/gin"

func InitRouter(Router *gin.RouterGroup) {
	InitTableRouter(Router)
	InitUserRouter(Router)

}
