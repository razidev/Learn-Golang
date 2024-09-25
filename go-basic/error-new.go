package main

import (
	"errors"
	"fmt"
	"log"
)

var errorMessage = errors.New("didn't accept the negative number")

func main() {
	fmt.Printf("%T\n", errorMessage)

	_, err := ErrorFunction(-10)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("End of File")
}

func ErrorFunction(i float64) (float64, error) {
	if i < 0 {
		// return 0, errorMessage
		return 0, fmt.Errorf("didn't accept the negative number: %v", i)
	}
	return 22, nil
}
