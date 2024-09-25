package main

import "fmt"

func main() {
	const firstName string = "Razi"
	const lastName = "Syahputro"
	const (
		familyName = "Syahputro"
		middleName = "Aziz"
	)

	// error
	// firstname = "Raz"
	// lastname = "Aziz"

	fmt.Println(firstName, middleName, lastName, familyName)
}
