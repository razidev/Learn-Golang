package main

import "fmt"

func main() {
	channel := make(chan string, 3)
	defer close(channel)

	channel <- "Iron Man"
	channel <- "Hulk"

	fmt.Println(<-channel)
	fmt.Println(<-channel)

	//won't deadlock bcs it's stored in buffered
	channel <- "Batman"
	channel <- "Nick Fury"

	fmt.Println(len(channel))
	fmt.Println(cap(channel))

	fmt.Println("EOF")
}
