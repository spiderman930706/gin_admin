package global

import (
	"github.com/spiderman930706/gin_admin/config"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

type Field struct {
	List   string
	Type   string //目前只想到password这种用途
	Schema *schema.Field
}

var (
	Config config.Config
	DB     *gorm.DB
	Tables map[string]map[string]*Field
)
