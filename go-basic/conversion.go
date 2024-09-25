package main

import "fmt"

func main() {
	var nilai32 int32 = 129
	var nilai64 int64 = int64(nilai32)
	var nilai8 int8 = int8(nilai32) //not enough space for int8 bcs the range from -128 to 127

	fmt.Println(nilai32)
	fmt.Println(nilai64)
	fmt.Println(nilai8)

	var name string = "Razi"
	var c byte = name[1]
	var get2ndChar string = string(c)

	fmt.Println(name)
	fmt.Println(c)
	fmt.Println(get2ndChar)
}
