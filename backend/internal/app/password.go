package app

import "golang.org/x/crypto/bcrypt"

func hashPassword(password string) (string, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashed), nil
}

func passwordsMatch(hashedPassword, givenPassword string) bool {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(givenPassword)) == nil
}
