package utils

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

func JWTSecret() string {
	return os.Getenv("JWT_SECRET")
}

func GenerateJWT(userID uint) (string, error) {
	claims := jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(JWTSecret()))
}
