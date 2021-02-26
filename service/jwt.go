package service

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/spiderman930706/gin_admin/global"
	"github.com/spiderman930706/gin_admin/models"
	"time"
)

func GenerateToken(user *models.User) (string, error) {
	var jwtSecret = []byte(global.Config.JWT.SigningKey)
	nowTime := time.Now()
	expireTime := nowTime.Add(time.Duration(global.Config.JWT.ExpireSecond) * time.Second)

	claims := models.CustomClaims{
		ID:       user.ID,
		Username: user.Username,
		IsAdmin:  user.IsAdmin,
		IsStaff:  user.IsStaff,
		RoleID:   user.Role.ID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "gin-blog",
		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(jwtSecret)

	return token, err
}

func ParseToken(token string) (*models.CustomClaims, error) {
	var jwtSecret = []byte(global.Config.JWT.SigningKey)
	tokenClaims, err := jwt.ParseWithClaims(token, &models.CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*models.CustomClaims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}

	return nil, err
}
