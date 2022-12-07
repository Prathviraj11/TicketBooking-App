package main

import (
	"fmt"
	"sync"
	"time"
)

var conferenceName = "Go Conference"
var remainingTickets uint = 50

// var bookings = make([]map[string]string,0)
var bookings = make([]userData, 0)

const confernceTickets = 50

type userData struct {
	firstName       string
	lastName        string
	email           string
	numberOfTickets uint
}

var wg = sync.WaitGroup{}

func main() {

	greetUser()

	for remainingTickets > 0 && len(bookings) <= 50 {

		firstName, lastName, email, userTickets := getUserInputs()
		isValidName, isValidEmail, isValidUsertickets := ValidateUserInputs(firstName, lastName, email, userTickets, remainingTickets)

		if isValidName && isValidEmail && isValidUsertickets {
			bookTickets(userTickets, firstName, lastName, email)

			wg.Add(1)
			go sendTickets(userTickets, firstName, lastName, email)

			firstNames := getFirstNames()
			fmt.Printf("The firstnames of bokings %v\n\n", firstNames)

			if remainingTickets == 0 {
				//end program
				fmt.Println("Our conference tickets are bookied out. Come back next year.")
				break
			}
		} else {
			if !isValidName {
				fmt.Println("FirstName or LastName you entered is wrong")
			}
			if !isValidEmail {
				fmt.Println("Email you entered is wrong")
			}
			if !isValidUsertickets {
				fmt.Println("Number of tikcets you entered is wrong")
			}
		}
	}
	wg.Wait()
}

func greetUser() {
	fmt.Printf("Welcome to %v\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are remaining\n", confernceTickets, remainingTickets)
	fmt.Printf("Get your tickets here\n\n")
}

func getFirstNames() []string {
	firstNames := []string{}
	for _, booking := range bookings {

		firstNames = append(firstNames, booking.firstName)
	}
	return firstNames
}

func getUserInputs() (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var userTickets uint

	fmt.Println("\nEnter your first name")
	fmt.Scan(&firstName)

	fmt.Println("Enter your last name")
	fmt.Scan(&lastName)

	fmt.Println("Enter your email address")
	fmt.Scan(&email)

	fmt.Println("Enter numbers of tickets")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets
}

func bookTickets(userTickets uint, firstName string, lastName string, email string) {
	remainingTickets = remainingTickets - userTickets

	//creating maps
	// var userData = make(map[string]string)
	// userData["firsName"] = firstName
	// userData["lastName"] = lastName
	// userData["email"] = email
	// userData["numberOfTickets"] = strconv.FormatUint(uint64(userTickets),10)

	//creating struct
	var userData = userData{
		firstName:       firstName,
		lastName:        lastName,
		email:           email,
		numberOfTickets: userTickets,
	}
	bookings = append(bookings, userData)

	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v\n\n", remainingTickets, conferenceName)

	//fmt.Printf("List of Bookings \n %v \n", bookings)
}

func sendTickets(userTickets uint, firstName string, lastName string, email string) {
	time.Sleep(50 * time.Second)
	var ticket = fmt.Sprintf("%v tikets for %v %v\n", userTickets, firstName, lastName)
	fmt.Println("\n#######################")
	fmt.Printf("Sending ticket %v to %v\n", ticket, email)
	fmt.Printf("#######################\n")
	wg.Done()
}
