package main

import (
	"fmt"
	"strconv"
)

func main() {
	channel := make(chan string)

	go func() {
		for i := 0; i < 10; i++ {
			channel <- "loops of " + strconv.Itoa(i)
		}
		close(channel)
	}()

	for data := range channel {
		fmt.Println("Receive data", data)
	}

	fmt.Println("EOF")
}
