package main

import (
	"fmt"
	"time"
)

func showsHelloWolrd() {
	fmt.Println("Hello World")
}

func showsNumber(number int) {
	fmt.Println("Number of", number)
}

func moreGoroutine() {
	for i := 0; i < 100000; i++ {
		go showsNumber(i)
	}
}

func main() {
	go showsHelloWolrd()
	fmt.Println("shows here")
	moreGoroutine()
	time.Sleep(5 * time.Second)

}
