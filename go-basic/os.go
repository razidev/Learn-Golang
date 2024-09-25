package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args
	fmt.Println("Argument:")
	fmt.Println(args)
	fmt.Println(args[0])

	hostname, err := os.Hostname()
	if err == nil {
		fmt.Println("hostname:", hostname)
	} else {
		fmt.Println("Error", err.Error())
	}

	username := os.Getenv("APP_USERNAME")
	fmt.Println(username)
}

/*
how to execute:
1. export APP_USERNAME=string
2. go run os.go args1 args2
**/
