package main

import (
	"fmt"
	"sync"
)

var counter = 0

func OnlyOnce() {
	counter++
}

func main() {
	once := sync.Once{}
	group := sync.WaitGroup{}

	for i := 0; i < 100; i++ {
		go func() {
			group.Add(1)
			once.Do(OnlyOnce) //just run once
			group.Done()
		}()
	}

	group.Wait()
	fmt.Println("counter", counter)
}
