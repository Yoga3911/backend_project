package utils

import "golang.org/x/crypto/bcrypt"

func HashAndSalt(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(hash), nil
}

func ComparePassword(plain string, hash string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(plain))
}
