package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	x := 0
	var mutex sync.Mutex
	var wg sync.WaitGroup
	fmt.Println("CPU:", runtime.NumCPU())
	fmt.Println("Goroutines", runtime.NumGoroutine())

	const num = 100
	wg.Add(100000)

	for i := 1; i <= 1000; i++ {
		go func() {
			for j := 1; j <= num; j++ {
				mutex.Lock()
				runtime.Gosched()
				x++
				mutex.Unlock()
				wg.Done()
			}
		}()
	}
	wg.Wait()
	fmt.Println("CPU:", runtime.NumCPU())
	fmt.Println("Goroutines", runtime.NumGoroutine())
	fmt.Println("counter = ", x)
}
