package main

import ("time")
import "log"

var s = "seven"

type User struct {
	FirstName string
	LastName string
	PhoneNumber string
	Age int
	BirthDate time.Time
}

func main() {
	user := User{
		FirstName: "Carlos",
		LastName: "Rodriguez",
		PhoneNumber: "123456789",
		Age: 30,
		BirthDate: time.Now(),
	}	

	log.Println(user.FirstName, user.PhoneNumber)
}