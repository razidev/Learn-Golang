package main

import "fmt"

type person struct {
	first      string
	last       string
	favFlavors []string
}

func main() {
	p1 := person{
		first: "Razi",
		last:  "Syahputro",
		favFlavors: []string{
			"Hot",
			"Coke",
			"Chocolate",
		},
	}

	p2 := person{
		first: "Hiyah",
		last:  "Iza",
		favFlavors: []string{
			"strawberry",
			"vanilla",
			"capuccino",
		},
	}

	m := map[string]person{
		p1.last: p1,
		p2.last: p2,
	}

	for _, v := range m {
		fmt.Println(v.first, v.last)
		for i, val := range v.favFlavors {
			fmt.Println(i, val)
		}
	}
}
