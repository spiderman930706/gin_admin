package gin_admin

import (
	"gin-admin/config"
	"gin-admin/core"
)


func Register(config config.Config, dst ...interface{}) {
	db := core.MysqlInit(config.Mysql)
	core.MigrateMysqlTables(db, dst...)
}
