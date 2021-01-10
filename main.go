package gin_admin

import (
	"log"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/spiderman930706/gin_admin/config"
	"github.com/spiderman930706/gin_admin/core"
	"github.com/spiderman930706/gin_admin/global"
	"github.com/spiderman930706/gin_admin/routers"
	"gorm.io/gorm/schema"
)

func RegisterConfigAndRouter(config config.Config, Router *gin.RouterGroup) {
	global.Config = config
	global.DB = core.MysqlInit(global.Config.Mysql)
	//Router.Use(middleware.JWTAuth())
	routers.InitRouter(Router)
}

func RegisterTables(migrate bool, dst ...interface{}) {
	if migrate {
		core.MigrateMysqlTables(global.DB, dst...)
	}
	global.Tables = make(map[string]map[string]*global.Field)
	for _, n := range dst {
		model := global.DB.Model(n)
		if err := model.Statement.Parse(n); err != nil {
			log.Printf("MySQL启动异常 %s", err)
			os.Exit(0)
		}
		tableName := model.Statement.Table
		global.Tables[tableName] = make(map[string]*global.Field)
		//先用表名和字段名来做数据的增删改，但在这里要取出所有定义了admin的参数，比如admin:"list:id;type:int"
		schemaField := model.Statement.Schema.FieldsByName
		ParseTag(schemaField, tableName)
	}
}

func ParseTag(fields map[string]*schema.Field, tableName string) {
	for _, v := range fields {
		RecordTag(v, tableName)
	}
}

func RecordTag(v *schema.Field, tableName string) {
	tag := v.Tag
	fieldName := v.DBName
	admin := tag.Get("admin")
	arr := strings.Split(admin, ";")
	m := &global.Field{}
	m.Schema = v
	for _, n := range arr {
		rr := strings.Split(n, ":")
		switch rr[0] {
		case "list":
			m.List = rr[1]
		case "type":
			m.Type = rr[1]
		}
	}
	global.Tables[tableName][fieldName] = m
}
