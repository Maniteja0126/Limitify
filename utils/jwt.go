package utils

import (
	"limitify/config"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateJWT(email string, plan string) (string, error) {
	claims := jwt.MapClaims{
		"email": email,
		"plan":  plan,
		"exp":   time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(config.JWTSecret)
}

func VerifyJWT(tokenString string) (*jwt.Token, error) {
	return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return config.JWTSecret, nil
	})
}
