package api

import (
	"github.com/gin-gonic/gin"
	"github.com/spiderman930706/gin_admin/models"
	"github.com/spiderman930706/gin_admin/service"
)

func UserLogin(c *gin.Context) {
	var L models.Login
	_ = c.ShouldBindJSON(&L)
	if err := L.Verify(); err != nil {
		FailWithMessage(err.Error(), c)
		return
	}
	U := &models.User{Username: L.Username, Password: L.Password}
	if err, user := service.Login(U); err != nil {
		FailWithMessage("用户名不存在或者密码错误", c)
		return
	} else {
		if token, err := service.GenerateToken(user); err != nil {
			FailWithMessage(err.Error(), c)
			return
		} else {
			OkWithDetailed(&models.LoginResult{
				Token: token,
			}, "登录成功", c)
			return
		}
	}
}
