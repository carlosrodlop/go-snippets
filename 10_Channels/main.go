package main

import (
	"10_Channels/helpers"
	"log"
)

const numPool = 1000

func calculateValue(intChan chan int) {
	randomNumber := helpers.RandomNumber(numPool)
	intChan <- randomNumber
}

func main() {
	// Channel not acepts int values
	intChan := make(chan int)
	//Best Practice: Declaratio to close the channel WHEN it is not used anymore
	defer close(intChan)

	//Call routine
	go calculateValue(intChan)

	//Channel listener, when it has a value, it is assigned to num
	num := <-intChan

	log.Println(num)
}
