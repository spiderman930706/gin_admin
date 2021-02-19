package routers

import (
	"github.com/gin-gonic/gin"
)

func InitRouter(AuthGroup *gin.RouterGroup, PubGroup *gin.RouterGroup, AdminGroup *gin.RouterGroup) {
	InitTableRouter(AuthGroup)
	InitPubUserRouter(PubGroup)
	InitSuperUserRouter(AdminGroup)
}
