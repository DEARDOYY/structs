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

func main() {
	firstName := getUserData("please enter your first name: ")
	lastName := getUserData("please enter your last name: ")
	birthDate := getUserData("please enter your birth date (MM/DD/YYYY): ")

	outPutUserDetails(firstName, lastName, birthDate)
}

func outPutUserDetails(firstName, lastName, birthDate string) {
	fmt.Println(firstName, lastName, birthDate)
}

func getUserData(promptText string) string {
	fmt.Println(promptText)

	var value string
	fmt.Scan(&value)
	return value
}
