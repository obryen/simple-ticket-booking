package main

import (
	"booking-app/helper"
	"fmt"
	"sync"
	"time"
)

var cName = "Nairobi tech day"

const tickets = 50

var RemainingTickets uint = tickets

var bookings = make([]UserInfo, 0)

type UserInfo struct {
	firstName       string
	email           string
	numberOfTickets uint
}

var wg = sync.WaitGroup{}

func main() {

	greetPeeps()

	userName, userEmail, userTickets := fetchUserInput()

	isValidName, isValidEmail, isValidTicket := helper.ValidateUserInput(userName, userEmail, userTickets, RemainingTickets)

	if !isValidEmail {
		fmt.Print("Invalid email \n")
	}
	if !isValidName {
		fmt.Print("Invalid name \n")
	}

	if !isValidTicket {
		fmt.Printf("Invalid input , we only have %v tickets remaining \n", RemainingTickets)
	}
	if isValidEmail && isValidName && isValidTicket {

		bookTicket(userTickets, userEmail, userName)
		wg.Add(1)
		go sendTicket(userTickets, userName, userEmail)

		fmt.Printf("The first names of bookings are %v \n", printFirstNames())

		noTicketsRemaining := RemainingTickets == 0
		if noTicketsRemaining {
			fmt.Print("Our conference is fully booked")

		}
	}
	wg.Wait()
}

func bookTicket(userTickets uint, userEmail string, userName string) {
	RemainingTickets = RemainingTickets - userTickets

	// create a map for a user
	var userData = UserInfo{
		firstName:       userName,
		email:           userEmail,
		numberOfTickets: RemainingTickets,
	}

	bookings = append(bookings, userData)
	fmt.Printf("List of bookings is %v \n", bookings)

	fmt.Printf("User %v has bought %v tickets, you will recieve an email  at %v \n", userName, userTickets, userEmail)
	fmt.Printf("%v tickets are now remaining\n", RemainingTickets)
}

func greetPeeps() {
	fmt.Printf("Welcome to the %v booking app\n", cName)
	fmt.Printf("There are a total of %v these are remaining %v.\n", tickets, RemainingTickets)
	fmt.Println("get your tickets now")
}

func printFirstNames() []string {
	firstNames := []string{}
	for _, booking := range bookings {

		firstNames = append(firstNames, booking.firstName)

	}
	return firstNames
}

func fetchUserInput() (string, string, uint) {
	var userName string
	var userEmail string
	var userTickets uint

	fmt.Println("Please enter your name:")
	fmt.Scan(&userName)

	fmt.Println("Please enter your email:")
	fmt.Scan(&userEmail)

	fmt.Println("Enter number of tickets:")
	fmt.Scan(&userTickets)

	return userName, userEmail, userTickets
}

func sendTicket(ticketNumber uint, name string, email string) {
	time.Sleep(50 * time.Second)
	ticket := fmt.Sprintf("%v tickets for %v user \n", ticketNumber, name)
	fmt.Println("####################")
	fmt.Printf("Sending ticket %v to email address %v \n", ticket, email)
	fmt.Println("####################")
	wg.Done()
}
