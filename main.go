package gin_admin

import (
	"github.com/gin-gonic/gin"
	"github.com/spiderman930706/gin_admin/config"
	"github.com/spiderman930706/gin_admin/core"
	"github.com/spiderman930706/gin_admin/global"
	"github.com/spiderman930706/gin_admin/routers"
	"log"
	"os"
	"reflect"
	"strings"
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
	global.Tables = make(map[string]map[string]map[string]string)
	for _, n := range dst {
		model := global.DB.Model(n)
		if err := model.Statement.Parse(n); err != nil {
			log.Printf("MySQL启动异常 %s", err)
			os.Exit(0)
		}
		tableName := model.Statement.Table
		global.Tables[tableName] = make(map[string]map[string]string)
		//先用表名和字段名来做数据的增删改，但在这里要取出所有定义了admin的参数，比如admin:"list:id;type:int"
		val := reflect.ValueOf(n)
		ParseTag(val, tableName)
	}
}

func ParseTag(v reflect.Value, tableName string) {
	if v.CanInterface() {
		t := v.Type()
		switch v.Kind() {
		case reflect.Ptr:
			ParseTag(v.Elem(), tableName)
		case reflect.Struct:
			for i := 0; i < v.NumField(); i++ {
				f := v.Field(i)
				if f.Kind() == reflect.Struct || f.Kind() == reflect.Ptr {
					tag := t.Field(i).Tag
					if tag != "" {
						RecordTag(tag, tableName, t.Field(i).Name)
					} else {
						ParseTag(f, tableName)
					}
				} else {
					tag := t.Field(i).Tag
					if tag != "" {
						RecordTag(tag, tableName, t.Field(i).Name)
					}
				}
			}
		}
	}
}

func RecordTag(tag reflect.StructTag, tableName string, fieldName string) {
	admin := tag.Get("admin")
	arr := strings.Split(admin, ";")
	m := make(map[string]string)
	for _, n := range arr {
		rr := strings.Split(n, ":")
		switch rr[0] {
		case "list":
			m["list"] = rr[1]
		case "type":
			m["type"] = rr[1]
		}
	}
	global.Tables[tableName][fieldName] = m
}
