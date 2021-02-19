# gin_admin

目的是做一个类似于django框架自带的admin后台管理，将注册的表在后台都可以方便管理，这是后端接口部分，前端有大佬感兴趣的话可以联系我

还在开发中，只实现了预想中的小部分功能

## 简单使用
```go
package main

import (
	"fmt"
	"gorm.io/gorm"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/spiderman930706/gin_admin"
	"github.com/spiderman930706/gin_admin/config"
	"github.com/spiderman930706/gin_admin/global"
	"github.com/spiderman930706/gin_admin/models"
	"github.com/spiderman930706/gin_admin/util"
)

type Article struct {
	models.Model

	TagID      int    `json:"tag_id" gorm:"index"`
	Tag        Tag    `json:"tag"`
	UserID     int    `json:"user_id" gorm:"index"`
	User       User   `json:"user"`
	Title      string `json:"title"`
	Desc       string `json:"desc"`
	Content    string `json:"content"`
	CreatedBy  string `json:"created_by"`
	ModifiedBy string `json:"modified_by"`
	State      int    `json:"state"`
}
type User struct {
	models.User
	Phone string `gorm:"column:phone_num;type:varchar(15);unique_index" json:"phone" admin:"name:手机"`
}

type Tag struct {
	models.Model

	Name       string `json:"name"`
	CreatedBy  string `json:"created_by"`
	ModifiedBy string `json:"modified_by"`
	State      int    `json:"state"`
}

func main() {
	con := config.Config{
		Mysql: config.Mysql{
			DbName:   "blog",
			User:     "root",
			Password: "",
			Host:     "127.0.0.1",
		},
		JWT: config.JWT{
			SigningKey:   "example-key",
			ExpireSecond: 7 * 24 * 3600,
		},
	}
	r := gin.Default()

	//注册gin_admin
	group := r.Group("admin")
	if err := gin_admin.RegisterConfigAndRouter(con, group); err != nil {
		log.Println(err)
		os.Exit(0)
	}
	if err := gin_admin.RegisterTables(true, &Article{}, &Tag{}, &User{}); err != nil {
		log.Println(err)
		os.Exit(0)
	}
	fmt.Println(global.Tables)

	CreateAdmin(global.DB)

	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", 8000),
		Handler:        r,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	fmt.Println(1 << 3)
	log.Println(s.ListenAndServe().Error())
}

func CreateAdmin(db *gorm.DB) {
	admin := make(map[string]interface{})
	var username = "admin"
	admin["username"] = username
	admin["password"] = util.MD5V([]byte("admin321!"))
	admin["is_admin"] = true
	admin["is_staff"] = true
	data := map[string]interface{}{}
	db.Table("users").Where("username = ?", username).Take(&data)
	if len(data) == 0 {
		db.Table("users").Create(&admin)
	}
}


```
