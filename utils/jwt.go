package utils

import (
	"github.com/golang-jwt/jwt/v5"
	"time"
)

// jwt package installed from: go get -u github.com/golang-jwt/jwt/v5

const secretKey = "superSecret"

func GenerateToken(email string, userId int64) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email":   email,
		"user_id": userId,
		"exp":     time.Now().Add(time.Hour * 2).Unix(),
	})
	return token.SignedString(secretKey)
}
