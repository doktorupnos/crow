package app

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// NewJWT returns a signed jwt token to respond with to the client.
func NewJWT(
	secret, subject string,
	expiresInMinutes time.Duration,
) (signedToken string, err error) {
	// TODO: Make the expiration configurable or define some good standards.
	now := time.Now()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.RegisteredClaims{
		Issuer:    "crow",
		IssuedAt:  jwt.NewNumericDate(now.UTC()),
		ExpiresAt: jwt.NewNumericDate(now.Add(expiresInMinutes).UTC()),
		Subject:   subject,
	})

	return token.SignedString([]byte(secret))
}

// ParseJWT returns an error when the token has expired
func ParseJWT(secret, tokenString string) (*jwt.Token, error) {
	return jwt.ParseWithClaims(
		tokenString,
		&jwt.RegisteredClaims{},
		func(t *jwt.Token) (interface{}, error) {
			return []byte(secret), nil
		},
	)
}
