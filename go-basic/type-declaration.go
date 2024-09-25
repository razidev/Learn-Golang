package main

import "fmt"

func main() {
	type NoKTP string
	type Married bool

	var myKtp NoKTP = "123455677854"
	var marriage Married = false
	var contoh string = "222222"
	var contohKTP NoKTP = NoKTP(contoh)

	fmt.Println(myKtp)
	fmt.Println(marriage)
	fmt.Println(contohKTP)
}

// type declaration uses for initialize with the name you want
