package main

import "fmt"

func sumAll(title string, numbers ...int) int { //varargs must be in the end of parameters / at the right not middle or left
	total := 0
	fmt.Println("value", numbers, title)
	for _, val := range numbers {
		total += val
	}
	return total
}

func main() {
	total := sumAll("First", 10, 9, 5, 3)
	fmt.Println("total", total)

	slice := []int{10, 5, 0, 20}
	total = sumAll("Second", slice...)
	fmt.Println("total", total)
}
