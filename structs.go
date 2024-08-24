package main

import "fmt"

func main() {
	firstName := getUserData("please enter your first name: ")
	lastName := getUserData("please enter your last name: ")
	birthDate := getUserData("please enter your birth date (MM/DD/YYYY): ")

	fmt.Println(firstName, lastName, birthDate)
}

func getUserData(promptText string) string {
	fmt.Println(promptText)

	var value string
	fmt.Scan(&value)
	return value
}
