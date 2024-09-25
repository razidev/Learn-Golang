package main

import "fmt"

func main() {
	a := []string{"razi", "aziz", "syahputro"}
	b := []string{"Nurul", "Hiyah"}
	fmt.Println(a)
	fmt.Println(b)

	c := [][]string{a, b}
	fmt.Println(len(c))
	fmt.Println(cap(c))
	fmt.Println(c)

	for i, v := range c {
		fmt.Println("record: ", i)
		for j, val := range v {
			fmt.Printf("%v: %v\n", j, val)
		}
	}
}
