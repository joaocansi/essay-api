package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/joaocansi/essay-api/config"
)

func GenerateJwt(id uint) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": id,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	})

	tokenString, err := token.SignedString([]byte(config.Env.JwtSecret))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
