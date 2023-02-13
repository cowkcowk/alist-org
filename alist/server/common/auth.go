package common

import (
	"time"

	"github.com/alist-org/alist/v3/internal/conf"
	"github.com/golang-jwt/jwt/v4"
)

var SecretKey []byte

type UserClaims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}

func GenerateToken(username string) (tokenString string, err error) {
	claim := UserClaims{
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(conf.Conf.TokenExpiresIn) * time.Hour)),
			IssuedAt: jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
		}}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	tokenString, err = token.SignedString(SecretKey)
	return tokenString, err
}