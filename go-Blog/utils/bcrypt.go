package utils

import "golang.org/x/crypto/bcrypt"

func PasswordCompare(hashed string, password []byte) (bool, error) {
	hashByte := []byte(hashed)
	err := bcrypt.CompareHashAndPassword(hashByte, password)

	if err != nil {
		return false, err
	}

	return true, nil
}

func PasswordHash(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 4)
	return string(bytes), err
}
