package cypher

import "golang.org/x/crypto/bcrypt"

func HashPassword(p []byte) (string, error) {
	hash, err := bcrypt.GenerateFromPassword(p, bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(hash), nil
}

func ValidatePassword(p []byte, hash []byte) bool {
	err := bcrypt.CompareHashAndPassword(hash, p)
	if err != nil {
		return false
	}

	return true
}
