package utils

import (
	"project/online-store/entity"
	"time"

	"github.com/golang-jwt/jwt"
)

type JwtTokenClaims struct {
	Email string      `json:"email"`
	Role  entity.Role `json:"role"`
	jwt.StandardClaims
}

func GenerateAccessToken(email, secret string, role entity.Role) (string, error) {
	claims := jwt.MapClaims{}
	claims["email"] = email
	claims["role"] = role
	claims["exp"] = time.Now().Add(time.Hour * 1).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secret))
}