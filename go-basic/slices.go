package main

import (
	"fmt"
	"slices"
)

func main() {
	names := []string{"Razi", "Aziz", "Syahputro"}
	values := []int{5, 1, 2, 3, 4}

	fmt.Println(slices.Min(names))
	fmt.Println(slices.Min(values))
	fmt.Println(slices.Max(names))
	fmt.Println(slices.Max(values))
	fmt.Println(slices.Contains(names, "Aziz"))
	fmt.Println(slices.Index(names, "Aziz"))
}
