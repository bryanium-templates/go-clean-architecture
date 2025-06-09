package utils

import (
	"os"
	"time"

	"github.com/bryanium-templates/go-clean-architecture/models"
	"github.com/golang-jwt/jwt"
)

func GenerateJWT(user *models.User) (string, error) {

	claims := jwt.MapClaims{
		"userId": user.ID.String(),
		"email": user.Email,
		"username": user.Username,
		"exp": time.Now().Add(time.Hour * 4).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	secret := os.Getenv("JWT_SECRET")


	tokenStr, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}

	return tokenStr, nil
}