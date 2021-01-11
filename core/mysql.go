package core

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"

	"github.com/spiderman930706/gin_admin/config"
)

func MysqlInit(m config.Mysql) (db *gorm.DB, err error) {
	dbName := m.DbName
	user := m.User
	password := m.Password
	host := m.Host
	dsn := user + ":" + password + "@tcp(" + host + ")/" + dbName + "?charset=utf8mb4&parseTime=True&loc=Local"
	mysqlConfig := mysql.Config{
		DSN:                       dsn,   // DSN data source name
		DefaultStringSize:         191,   // string 类型字段的默认长度
		DisableDatetimePrecision:  true,  // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,  // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,  // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false, // 根据版本自动配置
	}
	if db, err = gorm.Open(mysql.New(mysqlConfig), &gorm.Config{}); err != nil {
		log.Printf("MySQL启动异常 %s", err)
		return nil, err
	} else {
		sqlDB, _ := db.DB()
		sqlDB.SetMaxIdleConns(10)
		sqlDB.SetMaxOpenConns(100)
	}
	return db, nil
}

//数据表迁移
func MigrateMysqlTables(db *gorm.DB, dst ...interface{}) error {
	err := db.AutoMigrate(dst...)
	if err != nil {
		return err
	}
	log.Println("migrate table success")
	return nil
}
