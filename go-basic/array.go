package main

import "fmt"

// The capacity of the array cannot increase / decrease

func main() {
	// var names [...]string : will error It must be declared like line 27
	var names [3]string
	names[0] = "razi"
	names[1] = "aziz"
	names[2] = "syahputro"
	fmt.Println(names)

	var values = [3]int{
		22,
		04,
		1996,
	}

	fmt.Println(values[1])

	var checkLength [10]int
	fmt.Println(len(checkLength))
	fmt.Println(checkLength)

	var array = [...]int{
		1, 3, 5, 7, 9,
	}

	fmt.Println(cap(array))

}
