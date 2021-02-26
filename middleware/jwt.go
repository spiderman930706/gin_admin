package middleware

import (
	"github.com/spiderman930706/gin_admin/service"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/spiderman930706/gin_admin/api"
)

func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		// jwt鉴权取头部信息 x-token，postman使用api key认证
		token := c.Request.Header.Get("x-token")
		if token == "" {
			api.FailWithDetailed(gin.H{"reload": true}, "未登录或非法访问", c)
			c.Abort()
			return
		}
		claims, err := service.ParseToken(token)
		if err != nil {
			api.FailWithDetailed(gin.H{"reload": true}, err.Error(), c)
			c.Abort()
			return
		} else if time.Now().Unix() > claims.ExpiresAt {
			api.FailWithMessage("登录信息已失效", c)
			c.Abort()
			return
		}
		c.Set("claims", claims)
		c.Next()
	}
}
