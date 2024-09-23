package utils

import (
	"time"

	"github.com/golang-jwt/jwt"
)

var secretKey = "admin_secret"

func GenerateToken(adminID int) (string, error) {
	claims := jwt.MapClaims{
		"user_id": adminID,
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secretKey))
}
