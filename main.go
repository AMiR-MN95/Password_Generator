package main

import (
	"fmt"
	"math/rand"
	"regexp"
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

	for i := range password {
		password[i] = allChars[rand.Intn(len(allChars))]
	}

	return string(password)
}

func passwordMeetsCriteria(password string) bool {
	// Define regular expressions to match criteria
	containsLower := regexp.MustCompile(`[a-z]`)
	containsUpper := regexp.MustCompile(`[A-Z]`)
	containsNumber := regexp.MustCompile(`[0-9]`)
	containsSpecial := regexp.MustCompile(`[!@#%^&*()-_=+\[\]{}|;:'",.<>/?]`)

	return containsLower.MatchString(password) &&
		containsUpper.MatchString(password) &&
		containsNumber.MatchString(password) &&
		containsSpecial.MatchString(password)
}

func main() {
	var length int
	fmt.Print("Enter the desired password length: ")
	_, err := fmt.Scanf("%d", &length)
	if err != nil {
		fmt.Println("Invalid input. Please enter a valid integer.")
		return
	}

	password := generatePassword(length)

	fmt.Println("Generated Password:", password)

	if passwordMeetsCriteria(password) {
		fmt.Println("Password meets criteria.")
	} else {
		fmt.Println("Password does not meet criteria.")
	}
}
