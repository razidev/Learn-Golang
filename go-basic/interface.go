package main

import "fmt"

type HasName interface {
	GetName() string
}

func SayHello(hasName HasName) {
	fmt.Println("Hello", hasName.GetName())
}

type Person struct {
	Name string
}

type Animal struct {
	Name string
}

func (person Person) GetName() string {
	return person.Name
}

func (animal Animal) GetName() string {
	return animal.Name
}
func main() {
	razi := Person{
		Name: "Razi",
	}

	var cat Animal
	cat.Name = "Cito"

	SayHello(razi)
	SayHello(cat)
}
