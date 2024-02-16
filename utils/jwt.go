package utils

import (
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"strconv"
	"time"
)

// jwt package installed from: go get -u github.com/golang-jwt/jwt/v5

const secretKey = "superSecret"

func GenerateToken(userId int64, email string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": strconv.FormatInt(userId, 10),
		"email":   email,
		"exp":     time.Now().Add(time.Hour * 2).Unix(),
	})
	// the correct type by signing method chosen:
	// https://golang-jwt.github.io/jwt/usage/signing_methods/#signing-methods-and-key-types
	return token.SignedString([]byte(secretKey))
}

func VerifyToken(tokenString string) error {
	parsedToken, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// syntax to check the type assertion: type.(anotherType)
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		if token.Method.Alg() != jwt.SigningMethodHS256.Alg() {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		expirationTime, err := token.Claims.GetExpirationTime()
		if err != nil {
			return nil, err
		}
		if expirationTime == nil {
			return nil, errors.New("token is invalid. It does not have expiration key")
		}
		if time.Now().UTC().After(expirationTime.Time.UTC()) {
			return nil, errors.New("the token has expired")
		}
		return []byte(secretKey), nil
	})

	if err != nil {
		return err
	}

	if !parsedToken.Valid {
		return errors.New("invalid token")
	}

	claims, ok := parsedToken.Claims.(jwt.MapClaims)
	if !ok {
		return errors.New("invalid token claims type")
	}

	// we could make any database perform with these user data
	// try to get the value of type string
	email := claims["email"].(string)
	userIdString := claims["user_id"].(string)
	userId, _ := strconv.ParseInt(userIdString, 10, 64)
	fmt.Printf("Token verification. Email: %v, User_ID: %v\n", email, userId)
	return nil
}
