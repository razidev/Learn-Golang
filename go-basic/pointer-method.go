package main

import "fmt"

type Man struct {
	Name string
}

// func (man Man) Married() { it won't change the name
func (man *Man) Married() {
	man.Name = "Mr. " + man.Name
}

func main() {
	razi := Man{"Razi"}
	razi.Married()

	fmt.Println(razi)
}
