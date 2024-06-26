package models

import "github.com/dgrijalva/jwt-go"

type CustomClaims struct {
	ID          int
	NickName    string
	AuthorityId int
	jwt.StandardClaims
}
