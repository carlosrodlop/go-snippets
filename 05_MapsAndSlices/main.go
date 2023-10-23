package main

import (
	"log"
	"sort"
)

type User struct {
	FirstName string
	LastName  string
}

func main() {
	// maps
	myDogMap := make(map[string]string)
	myUserMap := make(map[string]User)

	myDogMap["dog1"] = "Brandy"
	myDogMap["dog2"] = "Dali"

	me := User{
		FirstName: "Carlos",
		LastName:  "Rodriguez",
	}

	myUserMap["me"] = me

	log.Println(myDogMap["dog2"])
	log.Println(myUserMap["me"].FirstName)

	// slices
	var mySlice []int

	mySlice = append(mySlice, 2)
	mySlice = append(mySlice, 1)
	mySlice = append(mySlice, 3)
	
	log.Println("Not order", mySlice)

	sort.Ints(mySlice)

	log.Println("In order", mySlice)

	// shorthand for slices

	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	log.Println(numbers)

	// print first two elements of slice, starting at first position
	log.Println(numbers[0:2])

	// create a slice of strings
	names := []string{"one", "seven", "fish", "cat"}
	log.Println(names)
}
