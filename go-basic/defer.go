package main

import "fmt"

func logging() {
	fmt.Println("selesai memanggil")
}

func runApplication(val int) {
	defer logging() //will be execute no matter error is show / not
	fmt.Println("Run App")
	result := 10 / val
	fmt.Println("res", result)
}
func main() {
	runApplication(10)
}
