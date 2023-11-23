package main

import "golang.org/x/crypto/bcrypt"

func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

func PasswordsMatch(hashedPassword, other string) bool {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(other)) == nil
}
