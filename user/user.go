package user

import (
	"errors"
	"fmt"
	"time"
)

// note : first character uppercase is can Outside package
// struct type
type User struct {
	firstName string
	lastName  string
	birthDate string
	createdAt time.Time
}

func (u User) OutPutUserDetails() {
	fmt.Println(u.firstName, u.lastName, u.birthDate)
}

func (u User) ClearUserName() {
	u.firstName = ""
	u.lastName = ""
}

func New(firstName, lastName, birthDate string) (*User, error) {
	if firstName == "" || lastName == "" || birthDate == "" {
		return nil, errors.New("first name, last name and birth date are require")
	}
	return &User{
		firstName: firstName,
		lastName:  lastName,
		birthDate: birthDate,
		createdAt: time.Now(),
	}, nil
}
