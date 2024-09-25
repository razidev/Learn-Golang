package main

import "fmt"

func getFullName() (string, string, string) {
	return "razi", "aziz", "syahputro"
}

func getFullName2() (firstName, middleName, lastName string) {
	firstName = "nurul"
	middleName = "iza"

	return
}

func main() {
	firstName, _, lastName := getFullName()
	fmt.Println("firstName", firstName)
	fmt.Println("lastName", lastName)

	a, _, c := getFullName2()
	fmt.Println("a", a)
	fmt.Println("c", c)
}
