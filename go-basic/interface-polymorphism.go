package main

import "fmt"

type person struct {
	first, last string
}

type marriage struct {
	person
	isMarried bool
}

func (m marriage) speak() {
	var message string
	if m.isMarried {
		message = "already married"
	} else {
		message = "not married yet"
	}
	fmt.Println("I am", m.first, m.last, "and I am", message)
}

type human interface {
	speak()
}

func bar(h human) {
	fmt.Println("I was passed into bar func", h)
}

func main() {
	people := marriage{
		person: person{
			first: "Lorem",
			last:  "Ipsum",
		},
		isMarried: false,
	}
	otherPeople := marriage{
		person: person{
			first: "dolor",
			last:  "sit amet",
		},
		isMarried: true,
	}

	people.speak()
	bar(people)
	otherPeople.speak()
	bar(otherPeople)
}
