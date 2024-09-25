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
	ctx, cancel := context.WithCancel(parent)

	destination := helper.CreateCounter(ctx)

	fmt.Println("total goroutine", runtime.NumGoroutine())

	for n := range destination {
		fmt.Println("Counter", n)
		if n == 10 {
			break
		}
	}

	cancel()

	time.Sleep(2 * time.Second)
	fmt.Println("total goroutine", runtime.NumGoroutine())
}
