package main

import (
	"fmt"
	"time"
)

func main() {
	x := 0

	for i := 1; i <= 1000; i++ {
		go func() {
			for j := 1; j <= 100; j++ {
				x++
			}
		}()
	}

	time.Sleep(5 * time.Second)
	fmt.Println("counter = ", x)
	//the result that should be 100000, the solution using mutex or atomic
}
