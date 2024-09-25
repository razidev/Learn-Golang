package main

import "fmt"

func getGoodBye(name string) string {
	return "Good bye " + name
}

func main() {
	sayGoodbye := getGoodBye
	fmt.Println(sayGoodbye("Razi"))
	fmt.Println(getGoodBye("Aziz"))
}
