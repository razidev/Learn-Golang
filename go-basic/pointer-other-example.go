package main

import "fmt"

func main() {
	a := 22
	fmt.Println("value of a:", a)
	fmt.Println("address of a:", &a)
	fmt.Printf("type value of a: %T\n", a)
	fmt.Printf("type address of a: %T\n", &a)

	var b *int = &a
	fmt.Println("address of b:", b)
	fmt.Println("value of b:", *b)

	*b = 90
	fmt.Println("value of a & b", a, *b)
	fmt.Println("adress of a & b", &a, b)

}
