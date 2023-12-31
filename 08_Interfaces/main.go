package main

import "fmt"

// Animal defines the INTERFACE for type Animal
// with a define SET OF FUCTIONS TO IMPLEMENT
type Animal interface {
	Says() string
	NumberOfLegs() int
}

func PrintInfo(a Animal) {
	fmt.Println("This animal says", a.Says(), "and has", a.NumberOfLegs(), "legs")
}

// Dog defines the dog type
type Dog struct {
	Name  string
	Breed string
}

// Says has a receiver of type *Dog, so it satisfies part of the interface requirements for Animal
// for the Dog type
func (d *Dog) Says() string {
	return "Woof"
}

// NumberOfLegs satisfies the rest of the Animal interface requirements for the Dog type
func (d *Dog) NumberOfLegs() int {
	return 4
}

// Gorilla defines the Gorilla type
type Gorilla struct {
	Name          string
	Color         string
	NumberOfTeeth int
}

// Says has a receiver of type *Gorilla, so it satisfies part of the interface requirements for Animal
// for the Gorilla type
func (d *Gorilla) Says() string {
	return "Ugh"
}

// NumberOfLegs satisfies the rest of the Animal interface requirements for the Gorilla type
func (d *Gorilla) NumberOfLegs() int {
	return 2
}

func main() {
	dog := Dog{
		Name:  "Samson",
		Breed: "German Shepherd",
	}

	// We can pass dog to PrintInfo(), since the Dog type implements the Animal interface by having all of the
	// necessary functions. The parameter is passed as a pointer since the functions for the type have pointer
	// receivers (which is the norm. See https://tour.golang.org/methods/4 and
	// https://tour.golang.org/methods/8 for more details).
	PrintInfo(&dog)

	gorilla := Gorilla{
		Name:          "Jock",
		Color:         "grey",
		NumberOfTeeth: 38,
	}

	// We can also pass gorilla to PrintInfo(), since the Gorilla type implements the Animal interface by having all of the
	// necessary functions. The parameter is passed as a pointer since the functions for the type have pointer
	// receivers (which is the norm. See https://tour.golang.org/methods/4 and
	// https://tour.golang.org/methods/8 for more details).
	PrintInfo(&gorilla)
}
