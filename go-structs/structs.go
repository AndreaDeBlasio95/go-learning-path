package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

type user struct {
	firstName string
	lastName  string
	age       string
	createdAt time.Time
}

func main() {
	userFirstName := getUserData("Please enter your first name: ")
	userLastName := getUserData("Please enter your last name: ")
	userAge := getUserData("Please enter your age (MM/DD/YYYY): ")

	// Create a new user
	appUser := user{
		firstName: userFirstName,
		lastName:  userLastName,
		age:       userAge,
		createdAt: time.Now(),
	}

	outputUserDetails(appUser)
}

func outputUserDetails(u user) {
	fmt.Printf("Name: %s %s\n", u.firstName, u.lastName)
	fmt.Printf("Age: %s\n", u.age)
	fmt.Printf("Created At: %s\n", u.createdAt.Format(time.RFC1123))
}

func getUserData(prompt string) string {
	fmt.Print(prompt)
	reader := bufio.NewReader(os.Stdin)
	value, _ := reader.ReadString('\n')
	return value[:len(value)-1] // Remove the trailing newline character
}
