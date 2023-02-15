package helper

import (
	"strings"
)

func ValidateUserInput(userName string, userEmail string, userTickets uint, ticketsRemaining uint) (bool, bool, bool) {
	isValidName := len(userName) >= 2 && len(userEmail) >= 2
	isValidEmail := strings.Contains(userEmail, "@")
	isValidTicket := userTickets > 0 && userTickets <= ticketsRemaining

	return isValidName, isValidEmail, isValidTicket
}
