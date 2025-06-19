package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var jwtSecret = []byte("your-secret-key")

func GenerateJWT(employeeID int, position string) (string, error) {
	claims := jwt.MapClaims{
		"employee_id": employeeID,
		"position":    position, // 예: "관리직", "운송직"
		"exp":         time.Now().Add(time.Hour * 24).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}
