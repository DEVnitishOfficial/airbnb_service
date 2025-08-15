package utils

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func HashedPassword(plainPassword string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(plainPassword), bcrypt.DefaultCost)

	if err != nil {
		fmt.Println("Error in hashing password", err)
		return "", err
	}
	return string(hash), nil
}

func CheckPasswordHash(plainPassword string, hashedPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(plainPassword))
	if err != nil {
		fmt.Println("Error in comparing password", err)
		return false
	}
	return true
}
