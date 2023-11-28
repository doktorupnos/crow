package jwt

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// Create creates a signed JWT.
func Create(secret, userID string, expiresIn time.Duration) (signedToken string, err error) {
	now := time.Now()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.RegisteredClaims{
		Issuer:    "crow",
		IssuedAt:  jwt.NewNumericDate(now),
		ExpiresAt: jwt.NewNumericDate(now.Add(expiresIn)),
		Subject:   userID,
	})
	return token.SignedString([]byte(secret))
}

// Parse returns an error if the token has expired.
func Parse(secret, tokenString string) (*jwt.Token, error) {
	return jwt.ParseWithClaims(
		tokenString,
		&jwt.RegisteredClaims{},
		func(t *jwt.Token) (interface{}, error) {
			return []byte(secret), nil
		},
	)
}
