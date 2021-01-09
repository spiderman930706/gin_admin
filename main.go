package gin_admin

import (
	"github.com/gin-gonic/gin"
	"github.com/spiderman930706/gin_admin/config"
	"github.com/spiderman930706/gin_admin/core"
	"github.com/spiderman930706/gin_admin/middleware"
	"github.com/spiderman930706/gin_admin/routers"
	"gorm.io/gorm"
)

var (
	Config config.Config
	DB     *gorm.DB
)

func Register(config config.Config, Router *gin.RouterGroup) {
	Config = config
	DB = core.MysqlInit(Config.Mysql)
	Router.Use(middleware.JWTAuth())
	routers.InitRouter(Router)
}

func MigrateTables(dst ...interface{}) {
	core.MigrateMysqlTables(DB, dst...)
	DB.Scopes()
}
