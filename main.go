package gin_admin

import (
	"github.com/spiderman930706/gin_admin/models"
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

func RegisterTables(migrate bool, dst ...models.AdminOperation) {
	if migrate {
		var res []interface{}
		for _, n := range dst {
			res = append(res, n)
		}
		core.MigrateMysqlTables(global.DB, res...)
	}
	global.Tables = make(map[string]*global.Table)
	for _, n := range dst {
		ParseSchema(n)
	}
}

func ParseSchema(n models.AdminOperation) {
	model := global.DB.Model(n)
	if err := model.Statement.Parse(n); err != nil {
		log.Printf("MySQL启动异常 %s", err)
		os.Exit(0)
	}
	tableName := model.Statement.Table
	//先用表名和字段名来做数据的增删改，但在这里要取出所有定义了admin的参数，比如admin:"list:id;type:int"
	schemaField := model.Statement.Schema.FieldsByName
	newFields := make(map[string]*global.Field)
	for _, v := range schemaField {
		fieldName, field := RecordTag(v, tableName)
		newFields[fieldName] = field
	}
	global.Tables[tableName] = &global.Table{
		Field:     newFields,
		CanDelete: n.CanDelete(),
		CanModify: n.CanModify(),
	}
}

func RecordTag(v *schema.Field, tableName string) (fieldName string, m *global.Field) {
	tag := v.Tag
	fieldName = v.DBName
	admin := tag.Get("admin")
	arr := strings.Split(admin, ";")
	m = &global.Field{}
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
	return
}
