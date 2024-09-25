package main

import "fmt"

func main() {
	x := foobar()
	fmt.Printf("%T\n", x)
	fmt.Println(x)
	fmt.Println(x())
}

func foobar() func() int {
	return func() int {
		return 22
	}
}
