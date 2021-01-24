package models

import "github.com/dgrijalva/jwt-go"

type CustomClaims struct {
	ID       uint
	Username string
	IsAdmin  bool
	IsStaff  bool
	RoleID   uint
	jwt.StandardClaims
}
