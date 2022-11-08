package criptography

import (
	"golang.org/x/crypto/bcrypt"
)

func EncodePassword(password string) string {
	passwordBytes := []byte(password)

	// Hashing the password with the default cost of 10
	hashedPassword, err := bcrypt.GenerateFromPassword(passwordBytes, bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}

	return string(hashedPassword)
}

func ComparePassword(password string, hashedPassword []byte) bool {
	passwordBytes := []byte(password)
	// Comparing the password with the hash
	err := bcrypt.CompareHashAndPassword(hashedPassword, passwordBytes)
	return err == nil
}
