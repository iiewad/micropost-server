package utils

import "golang.org/x/crypto/bcrypt"

// PasswordSecret *
func PasswordSecret(pw string) (string, error) {
	password, err := bcrypt.GenerateFromPassword([]byte(pw), bcrypt.DefaultCost)
	if err != nil {
		return string(password), err
	}
	return string(password), err
}
