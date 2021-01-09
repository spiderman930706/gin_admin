package gin_admin

import (
	"log"
	"os"

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
	global.Tables = make(map[string]interface{})
	for _, n := range dst {
		model := global.DB.Model(n)
		if err := model.Statement.Parse(n); err != nil {
			log.Printf("MySQL启动异常 %s", err)
			os.Exit(0)
		}
		global.Tables[model.Statement.Table] = n //先用表名和字段名来做数据的增删改，但在这里要取出所有定义了admin的参数，比如admin:"list:id;type:int"

	}
}
