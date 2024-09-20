package main

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func hashPassword(password string) (string, error) {

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

func checkPassword(hashedPassword, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}

func main() {
	password := "mySecretPassword123"

	hashedPassword, err := hashPassword(password)
	if err != nil {
		fmt.Println("error while hashing the password:", err)
		return
	}

	fmt.Println("hashed password:", hashedPassword)

	isValid := checkPassword(hashedPassword, "mySecretPassword123")
	fmt.Println("password is correct?", isValid)

	isValid = checkPassword(hashedPassword, "wrongPassword")
	fmt.Println("password is correct?", isValid)
}
