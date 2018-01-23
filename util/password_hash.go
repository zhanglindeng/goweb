package util

import "golang.org/x/crypto/bcrypt"

var PasswordHashCost = bcrypt.DefaultCost

func PasswordHash(s string) (string, error) {
	b, err := bcrypt.GenerateFromPassword([]byte(s), PasswordHashCost)
	return string(b), err
}
