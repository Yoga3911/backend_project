package utils

import "golang.org/x/crypto/bcrypt"

func hashAndSalt(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(hash), nil
}

func comparePassword(plain string, hash string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(plain))
}
