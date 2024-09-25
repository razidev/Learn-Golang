package main

import (
	"Udemy-PZN/helper"

	"context"
	"fmt"
	"runtime"
	"time"
)

func main() {
	fmt.Println("total goroutine", runtime.NumGoroutine())

	parent := context.Background()
	ctx, cancel := context.WithTimeout(parent, 5*time.Second)
	defer cancel()

	destination := helper.CreateCounter(ctx)
	fmt.Println("total goroutine", runtime.NumGoroutine())

	for n := range destination {
		fmt.Println("Counter", n)
	}

	time.Sleep(2 * time.Second)
	fmt.Println("total goroutine", runtime.NumGoroutine())
}
