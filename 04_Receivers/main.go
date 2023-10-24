package main

import (
	"fmt"
	"log"
	"math"
)

type myStruct struct {
	FirstName string
}

// Receivers are like "methods" for structs. They are functions that are attached to a struct.
func (m myStruct) printFirstName() string {
	return m.FirstName
}

type Vertex struct {
	X, Y float64
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	var myVar myStruct
	myVar.FirstName = "Carlos"

	myVar2 := myStruct{
		FirstName: "Mary",
	}

	log.Println("myVar is set to", myVar.FirstName)
	log.Println("myVar is set to", myVar.printFirstName())
	log.Println("myVar2 is set to", myVar2.FirstName)
	log.Println("myVar2 is set to", myVar2.printFirstName())

	v := Vertex{3, 4}
	v.Scale(10)
	fmt.Println(v.Abs())

}
