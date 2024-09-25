package main

import (
	"fmt"
	"strconv"
)

func main() {
	boolean, err := strconv.ParseBool("false")

	if err == nil {
		fmt.Println(boolean)
	} else {
		fmt.Println(err.Error())
	}

	number, err := strconv.ParseInt("1600000", 10, 64)
	if err == nil {
		fmt.Println(number)
	} else {
		fmt.Println(err.Error())
	}

	val := strconv.FormatInt(1000000, 10)
	fmt.Println(val)

	valueInt, _ := strconv.Atoi("2000000")
	fmt.Println(valueInt)
}
