package main

import (
	"Udemy-PZN/helper"
	"fmt"
	"time"
)

func main() {
	channel := make(chan string)
	defer close(channel)

	go helper.GiveMeResponse(channel, "Hulk")

	data := <-channel
	fmt.Println("data:", data)

	time.Sleep(5 * time.Second)
}
