package main

import (
	"flag"
	"fmt"
)

func main() {
	var host *string = flag.String("host", "localhost", "Put your database host")
	var user *string = flag.String("user", "root", "Put your database user")
	var number *int = flag.Int("number", 100, "Put your number")

	flag.Parse()

	fmt.Println(*host)
	fmt.Println(*user)
	fmt.Println(*number)
}

/*
how to execute:
1. go run flag.go -host=localhost -user=YourUser -number=90
// if not using flag, the value is default
**/
