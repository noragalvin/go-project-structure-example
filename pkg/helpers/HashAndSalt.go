package helpers

import (
	"golang.org/x/crypto/bcrypt"
)

// HashAndSalt create strong password with salt
func HashAndSalt(pwd string) (string, error) {
	byteString := []byte(pwd)
	hash, err := bcrypt.GenerateFromPassword(byteString, bcrypt.MinCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}
