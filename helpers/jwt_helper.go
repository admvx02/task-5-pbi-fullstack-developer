package helpers

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

var jwtKey = []byte("your-secret-key")

func GenerateToken(userId uint) (string, error) {
	claims := jwt.MapClaims{
		"userId": userId,
		"exp":    time.Now().Add(time.Hour * 24).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtKey)
}

func ValidateToken(tokenString string) (*jwt.Token, error) {
	return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
}
