package utils

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

var (
	JwtKey = []byte("_konnect_is_not_connect_") // Replace with your own secret key
)

type Claims struct {
	Email string `json:"email"`
	jwt.StandardClaims
}

func GenerateToken(email string) (string, error) {
	expirationTime := time.Now().Add(15 * time.Minute)

	claims := &Claims{
		Email: email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	jwtToken, err := token.SignedString(JwtKey)
	if err != nil {
		return "", err
	}

	return jwtToken, nil
}
