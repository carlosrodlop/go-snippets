package main

import (
	"09_Packages/helpers"
	"log"
)

func main() {
	var myVar helpers.SomeType
	myVar.TypeName = "Some name"
	log.Println(myVar.TypeName)
}
