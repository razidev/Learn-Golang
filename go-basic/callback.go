package main

import "fmt"

func main() {
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	total := sum(slice...)
	callback := odd(sum, slice...)
	fmt.Println("total number:", total)
	fmt.Println("odd number:", callback)
}

func sum(vari ...int) int {
	total := 0
	for _, v := range vari {
		total += v
	}
	return total
}

func odd(call func(x ...int) int, v ...int) int {
	var j []int
	for _, val := range v {
		if val%2 == 1 {
			j = append(j, val)
		}
	}

	return call(j...)
}
