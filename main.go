package main

import (
	"fmt"
	"strings"
	"strconv"
)

const conferenceName = "Go Conference"

const conferenceTickets int = 50
var remainingTickets uint = 50
var bookings = make([]map[string]string)


func main() {

	greetUsers()

	for {
		firstName, lastName, email, userTickets := getUserInput()
		isValidName, isValidEmail, isValidTicketNumber := ValidateUserInput(firstName, lastName, email, userTickets, remainingTickets)
	
		if isValidName && isValidEmail && isValidTicketNumber {
			bookTicket(userTickets, firstName, lastName, email)
			// call function print first names
			firstNames := getFirstNames()
			fmt.Printf("First names: %v\n", firstNames)

			if remainingTickets == 0 {
				//end program
				fmt.Println("Our conference is sold out")
				break
			}
		} else {
			if !isValidName {
				fmt.Println("First name or last name is too short")
			}
			if !isValidEmail {
				fmt.Println("Email doesnt have '@'")
			}
			if !isValidTicketNumber {
				fmt.Println("Number of tickets is invalid")
			}
		}
	}
}

func greetUsers() {
	fmt.Printf("Welcome to %v booking application!. We are selling %v tickets.\n", conferenceName, conferenceTickets)
	fmt.Printf("Get your ticket or ass ready. We still have %v available\n", remainingTickets)
}

func getFirstNames() []string {
	firstNames := []string{}
	for _, booking := range bookings {
		var names = strings.Fields(booking)
		firstNames = append(firstNames, names[0])
	}
fmt.Printf("All bookings: %v\n", bookings)
return firstNames
}

func getUserInput() (string, string, string, uint) {
	var firstName string
		var lastName string
		var email string
		var userTickets uint
		//Asking user for their name
		fmt.Println("Enter your first name")
		fmt.Scan(&firstName)

		fmt.Println("Enter your last name")
		fmt.Scan(&lastName)

		fmt.Println("Enter your email")
		fmt.Scan(&email)

		fmt.Println("Enter number of tickets:")
		fmt.Scan(&userTickets)
		return firstName, lastName, email, userTickets
}

func bookTicket(userTickets uint, firstName string, lastName string, email string) {
	remainingTickets = remainingTickets - userTickets
	
	//create a map for a user
	var userData = make([]map[string]string, 0)
	userData["firstName"] = firstName
	userData["lastName"] = lastName
	userData["email"] = email
	userData["numberOfTickets"] = strconv.FormatUint(uint64(userTickets), 10)

	bookings = append(bookings, firstName + " " + lastName)

	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)
}