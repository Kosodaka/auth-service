package models

import (
	"crypto/rand"
	"github.com/dgrijalva/jwt-go"
	"time"
)

type tokenClaims struct {
	Guid string `json:"guid"`
	jwt.StandardClaims
}

const (
	signingKey = "djkfgfguvhodvjpioposdffsdf"
)

func GenerateAccessToken(guid string) (string, error) {
	access := jwt.NewWithClaims(jwt.SigningMethodHS512, &tokenClaims{
		guid,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Minute * 15).Unix(),
		},
	})
	accessSigned, err := access.SignedString([]byte(signingKey))
	if err != nil {
		return "", err
	}

	return accessSigned, nil

}

func GenerateRefreshToken() (string, error) {
	refreshBytes := make([]byte, 32) // Произвольная длина
	_, err := rand.Read(refreshBytes)
	if err != nil {
		return "", err
	}

	return string(refreshBytes), nil
}
