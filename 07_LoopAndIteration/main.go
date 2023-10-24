package main

import "log"

// There are only for loops
func main() {

	for i := 0; i <= 5; i++ {
		log.Println(i)
	}

	animals := make(map[string]string)
	animals["dog"] = "woof"
	animals["cat"] = "meow"

	for k, v := range animals { // k is the key, v is the value
		log.Println(k, v)
	}

	type User struct {
		FirstName string
		LastName  string
		Email     string
		Age       int
	}

	var users []User
	users = append(users, User{"John", "Smith", "john@smith.com", 30})
	users = append(users, User{"Mary", "Jones", "mary@jones.com", 20})
	users = append(users, User{"Sally", "Brown", "sally@smith.com", 45})
	users = append(users, User{"Alex", "Anderson", "alex@smith.com", 17})

	for i, l := range users {
		log.Println(i, l.FirstName, l.LastName, l.Email, l.Age)
	}

	for _, l := range users { // _ is a blank identifier. Used when index (i) is not needed
		log.Println(l.FirstName, l.LastName, l.Email, l.Age)
	}

}
