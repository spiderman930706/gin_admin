package global

import (
	"github.com/spiderman930706/gin_admin/config"
	"gorm.io/gorm"
)

var (
	Config config.Config
	DB     *gorm.DB
	Tables map[string]map[string]map[string]string
)
