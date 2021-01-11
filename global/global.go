package global

import (
	"github.com/spiderman930706/gin_admin/config"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

type Table struct {
	Field     map[string]*Field
	CanModify bool
	CanDelete bool
}

type Field struct {
	List   string //列表展示页名称，同时用来判断是否在列表页展示该字段
	Type   string //字段类型，目前只想到password这种用途，其余使用gorm中的类型
	Schema *schema.Field
}

var (
	Config config.Config
	DB     *gorm.DB
	Tables map[string]*Table
)
