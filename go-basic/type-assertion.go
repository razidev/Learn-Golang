package main

import (
	"fmt"
)

func random() interface{} {
	return "Learn to fun"
}

func main() {
	var result interface{} = random()
	//if you wrong use data type it will be panic & stop the program
	// var resultString string = result.(string)
	// fmt.Println(resultString)

	switch val := result.(type) {
	case string:
		fmt.Println("value", val, "is string")
	case int:
		fmt.Println("value", val, "is int")
	default:
		fmt.Println("Unknown type")
	}
}
