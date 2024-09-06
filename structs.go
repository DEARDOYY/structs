package main

import (
	"fmt"

	"example.com/structs/user"
)

func main() {
	userFirstName := getUserData("please enter your first name: ")
	userLastName := getUserData("please enter your last name: ")
	userBirthDate := getUserData("please enter your birth date (MM/DD/YYYY): ")

	var appUser *user.User

	appUser, err := user.New(userFirstName, userLastName, userBirthDate)

	if err != nil {
		fmt.Println(err)
		return
	}

	admin := user.NewAdmin("test@example.com", "test123")

	admin.User.OutPutUserDetails()
	admin.User.OutPutUserDetails()

	appUser.OutPutUserDetails()
	appUser.ClearUserName()
}

func getUserData(promptText string) string {
	fmt.Print(promptText)

	var value string
	fmt.Scan(&value)
	return value
}
