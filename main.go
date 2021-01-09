package gin_admin

import (
	"github.com/gin-gonic/gin"
	"github.com/spiderman930706/gin_admin/config"
	"github.com/spiderman930706/gin_admin/core"
	"github.com/spiderman930706/gin_admin/global"
	"github.com/spiderman930706/gin_admin/middleware"
	"github.com/spiderman930706/gin_admin/routers"
)

func Register(config config.Config, Router *gin.RouterGroup) {
	global.Config = config
	global.DB = core.MysqlInit(global.Config.Mysql)
	Router.Use(middleware.JWTAuth())
	routers.InitRouter(Router)
}

func MigrateTables(dst ...interface{}) {
	core.MigrateMysqlTables(global.DB, dst...)
	global.DB.Scopes()
}
