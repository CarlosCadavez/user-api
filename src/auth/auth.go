package auth

import (
	"time"
	"user-api/src/config"

	jwt "github.com/dgrijalva/jwt-go"
)

func CreateToken(userId uint64) (string, error) {
	permission := jwt.MapClaims{}
	permission["authorized"] = true
	permission["exp"] = time.Now().Add(time.Hour * 6).Unix()
	permission["userId"] = userId

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, permission)

	return token.SignedString(config.SecretKey)
}
