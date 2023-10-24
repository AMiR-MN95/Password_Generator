package main

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	lowerChars   = "abcdefghijklmnopqrstuvwxyz"
	upperChars   = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	numberChars  = "0123456789"
	specialChars = "!@#$%^&*()-_=+[]{}|;:'\",.<>/?"
)

func generatePassword(length int) string {
	allChars := lowerChars + upperChars + numberChars + specialChars
	password := make([]byte, length)

	rand.Seed(time.Now().UnixNano())

	for i := range password {
		password[i] = allChars[rand.Intn(len(allChars))]
	}

	return string(password)
}

func main() {
	password := generatePassword(8)
	fmt.Println("Generated Password:", password)
}
