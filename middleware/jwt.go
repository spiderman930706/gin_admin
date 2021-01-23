package middleware

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"

	"github.com/spiderman930706/gin_admin/api"
	"github.com/spiderman930706/gin_admin/global"
	"github.com/spiderman930706/gin_admin/models"
)

type CustomClaims struct {
	ID       uint
	Username string
	IsAdmin  bool
	IsStaff  bool
	jwt.StandardClaims
}

func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		// jwt鉴权取头部信息 x-token，postman使用api key认证
		token := c.Request.Header.Get("x-token")
		if token == "" {
			api.FailWithDetailed(gin.H{"reload": true}, "未登录或非法访问", c)
			c.Abort()
			return
		}
		claims, err := ParseToken(token)
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

func GenerateToken(user models.User) (string, error) {
	var jwtSecret = []byte(global.Config.JWT.SigningKey)
	nowTime := time.Now()
	expireTime := nowTime.Add(time.Duration(global.Config.JWT.ExpireSecond) * time.Second)

	claims := CustomClaims{
		user.ID,
		user.Username,
		user.IsAdmin,
		user.IsStaff,
		jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "gin-blog",
		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(jwtSecret)

	return token, err
}

func ParseToken(token string) (*CustomClaims, error) {
	var jwtSecret = []byte(global.Config.JWT.SigningKey)
	tokenClaims, err := jwt.ParseWithClaims(token, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*CustomClaims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}

	return nil, err
}
