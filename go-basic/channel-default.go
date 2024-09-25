package main

import (
	"Udemy-PZN/helper"
	"fmt"
)

func main() {
	channel1 := make(chan string)
	channel2 := make(chan string)

	defer close(channel1)
	defer close(channel2)

	go helper.GiveMeResponse(channel1, "Peacemaker")
	go helper.GiveMeResponse(channel2, "The Flash")

	counter := 0
	for {
		select {
		case data := <-channel1:
			fmt.Println("Data dari channel 1", data)
			counter++
		case data := <-channel2:
			fmt.Println("Data dari channel 2", data)
			counter++
		default:
			fmt.Println("Menunggu data")
		}

		if counter == 2 {
			break
		}
	}
}
