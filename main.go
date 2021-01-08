package gin_admin

import (
	"gin_admin/config"
	"gin_admin/core"
	"gin_admin/middleware"
	"gin_admin/routers"
	"github.com/gin-gonic/gin"
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
}
