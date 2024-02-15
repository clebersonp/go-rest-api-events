package utils

import (
	"github.com/golang-jwt/jwt/v5"
	"time"
)

// jwt package installed from: go get -u github.com/golang-jwt/jwt/v5

const secretKey = "superSecret"

func GenerateToken(userId int64, email string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": userId,
		"email":   email,
		"exp":     time.Now().Add(time.Hour * 2).Unix(),
	})
	// the correct type by signing method chosen:
	// https://golang-jwt.github.io/jwt/usage/signing_methods/#signing-methods-and-key-types
	return token.SignedString([]byte(secretKey))
}
