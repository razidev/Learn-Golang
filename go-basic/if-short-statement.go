package main

import "fmt"

func main() {
	name := "razi aziz"

	if length := len(name); length > 5 {
		fmt.Println("nama benar", name)
	} else {
		fmt.Println("Nama terlalu pendek")
	}
}
