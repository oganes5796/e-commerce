package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var jwtKey = []byte("HHGFh78d9inaIY9ouslin&fus") // Секретный ключ для подписи токенов

// Генерация JWT токена
func GenerateJWT(userID int) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(time.Hour * 24).Unix(), // Токен действителен 24 часа
	})

	return token.SignedString(jwtKey)
}
