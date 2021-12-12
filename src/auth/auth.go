package auth

import (
	"errors"
	"fmt"
	"net/http"
	"strings"
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

func ValidateToken(r *http.Request) error {
	tokenAsString := extractToken(r)
	fmt.Println("Token ->>> " + tokenAsString)
	token, err := jwt.Parse(tokenAsString, getKey)
	if err != nil {
		return err
	}

	if _, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return nil
	}

	return errors.New("Invalid token")
}

func extractToken(r *http.Request) string {
	token := r.Header.Get("Authorization")
	if len(strings.Split(token, " ")) == 2 {
		return strings.Split(token, " ")[1]
	}

	return ""
}

func getKey(token *jwt.Token) (interface{}, error) {
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, fmt.Errorf("Unexpected sign method! %v", token.Header["alg"])
	}

	return config.SecretKey, nil
}
