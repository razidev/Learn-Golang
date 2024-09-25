package main

import "fmt"

func main() {
	name := "syahputrooo"

	switch name {
	case "razi":
		fmt.Println("hello")
		fmt.Println(name)
	case "aziz":
		fmt.Println("hai")
		fmt.Println(name)
	default:
		fmt.Println("name not found")
	}

	switch length := len(name); length > 5 {
	case true:
		fmt.Println(name, "terlalu panjang")
	case false:
		fmt.Println("yay cukup")
	}

	length := len(name)
	switch {
	case length > 10:
		fmt.Println(name, "kepanjangan")
		fallthrough
	case length > 6:
		fmt.Println(name, "panjang")
	default:
		fmt.Println("cukup")
	}
}
