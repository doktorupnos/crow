package jwt

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const issuer = "crow"

func Create(secret, userID string, expiresIn time.Duration) (signedToken string, err error) {
	now := time.Now()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.RegisteredClaims{
		Issuer:    issuer,
		IssuedAt:  jwt.NewNumericDate(now),
		ExpiresAt: jwt.NewNumericDate(now.Add(expiresIn)),
		Subject:   userID,
	})
	return token.SignedString([]byte(secret))
}

func Parse(secret, token string) (*jwt.Token, error) {
	return jwt.ParseWithClaims(token, &jwt.RegisteredClaims{}, func(t *jwt.Token) (any, error) {
		return []byte(secret), nil
	})
}
