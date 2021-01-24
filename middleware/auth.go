package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/spiderman930706/gin_admin/api"
	"github.com/spiderman930706/gin_admin/models"
)

func AdminAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		// jwt鉴权取头部信息 x-token，postman使用api key认证
		claims, exists := c.Get("claims")
		if !exists {
			api.FailWithMessage("未登录或非法访问", c)
			c.Abort()
			return
		}
		if claims, ok := claims.(*models.CustomClaims); !ok {
			api.FailWithMessage("访问出错", c)
			c.Abort()
			return
		} else {
			if !claims.IsStaff || !claims.IsAdmin {
				api.FailWithMessage("非法访问", c)
				c.Abort()
				return
			}
			if accessOk := AccessControl(claims); !accessOk {
				api.FailWithMessage("权限不足", c)
				c.Abort()
				return
			}
		}
		c.Next()
	}
}

func AccessControl(claims *models.CustomClaims) (ok bool) {
	// todo 在这里校验用户访问权限
	return
}
