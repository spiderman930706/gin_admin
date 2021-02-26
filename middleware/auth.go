package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/spiderman930706/gin_admin/api"
	"github.com/spiderman930706/gin_admin/global"
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
			if !claims.IsAdmin {
				api.FailWithMessage("权限不足", c)
				c.Abort()
				return
			}
		}
		c.Next()
	}
}

func RoleAuth() gin.HandlerFunc {
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
			if !claims.IsStaff {
				api.FailWithMessage("非法访问", c)
				c.Abort()
				return
			}
			if !claims.IsAdmin {
				if accessOk := accessControl(claims, c); !accessOk {
					api.FailWithMessage("权限不足", c)
					c.Abort()
					return
				}
			}
		}
		c.Next()
	}
}

func accessControl(claims *models.CustomClaims, c *gin.Context) (ok bool) {
	roleID := claims.RoleID
	if roleID == 0 {
		return
	}
	method := c.Request.Method
	table := c.Param("table")
	var role models.Role
	global.DB.Where("id = ?", roleID).Preload("Auth").First(&role)
	c.Set("role", &role)
	if table != "" {
		for _, auth := range role.Auth {
			if auth.Method == method && auth.TableName == table {
				return true
			}
		}
	} else {
		return true
	}
	return
}
