package helper

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
)

// function to hashed password
func HashPassword(password string) (string, error) {
	hashedPasswerd, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(hashedPasswerd), nil
}

// function to check password
func CheckPasword(hashedPassword, password string) (bool, error) {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil {
		return false, errors.New("password not match")
	}

	// password match
	return true, nil
}
