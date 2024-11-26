package auth

import (
	"ecommerce/config"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func CreateJwtToken(secret []byte, userID int) (string, error) {
	expiration := time.Second * time.Duration(config.Envs.JwtExpirationTime)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userId":    userID,
		"expiresAt": time.Now().Add(expiration).Unix(),
	})

	tokenString, err := token.SignedString(secret)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
