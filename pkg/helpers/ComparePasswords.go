package helpers

import (
	"golang.org/x/crypto/bcrypt"
)

// ComparePasswords ... So sánh password đã được mã hoá và password truyền vào
// hashedPwd là Password đã được mã hoá
// plainPwd là Password nguyên bản
func ComparePasswords(hashedPwd string, plainPwd string) bool {
	byteString := []byte(plainPwd)
	byteHash := []byte(hashedPwd)
	err := bcrypt.CompareHashAndPassword(byteHash, byteString)
	if err != nil {
		return false
	}
	return true
}
