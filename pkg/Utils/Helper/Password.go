package helper

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		fmt.Println("error inside hash password ")
		return "nil", fmt.Errorf("couldn't hash password %w ", err)
	}
	return string(hashPassword), nil
}

func ComparePassword(hashPassword, inputPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashPassword), []byte(inputPassword))
	if err != nil {
		return false
	}
	return true
}

