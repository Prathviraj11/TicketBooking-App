package main

import "strings"

func ValidateUserInputs(firstName string, lastname string, email string, usertickets uint,remainingTickets uint) (bool, bool, bool) {
	isValidName := len(firstName) >= 2 && len(lastname) >= 2
	isValidEmail := strings.Contains(email, "@")
	isValidUsertickets := usertickets > 0 && usertickets <= remainingTickets
	return isValidName, isValidEmail, isValidUsertickets
}