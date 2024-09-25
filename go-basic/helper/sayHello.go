package helper

import "fmt"

var version = "1.8.9"
var Application = "Belajar Golang"

func SayHello(name string) {
	fmt.Println("Hello", name)
}

func sayGoodBye(name string) {
	fmt.Println("Good bye", name)
}

//if variable or function started with Uppercase it means public, vice versa
