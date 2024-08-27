package main

import (
	"fmt"
	"time"
)

// note : first character uppercase is can Outside package
// struct type
type user struct {
	firstName string
	lastName  string
	birthDate string
	createdAt time.Time
}

func (u user) outPutUserDetails() {
	fmt.Println(u.firstName, u.lastName, u.birthDate)
}

func main() {
	userFirstName := getUserData("please enter your first name: ")
	userLastName := getUserData("please enter your last name: ")
	userBirthDate := getUserData("please enter your birth date (MM/DD/YYYY): ")

	var appUser user

	appUser = user{
		firstName: userFirstName,
		lastName:  userLastName,
		birthDate: userBirthDate,
		createdAt: time.Now(),
	}

	appUser.outPutUserDetails()
}

func getUserData(promptText string) string {
	fmt.Print(promptText)

	var value string
	fmt.Scan(&value)
	return value
}
