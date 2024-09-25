package main

import (
	"fmt"
	"sync"
	"time"
)

func Timer() {
	timer := time.NewTimer(5 * time.Second)
	fmt.Println("Timer Now", time.Now())

	time := <-timer.C
	fmt.Println("Timer Now After", time)
}

func TimerAfter() {
	channel := time.After(5 * time.Second)
	fmt.Println("Timer After Now", time.Now())

	time := <-channel
	fmt.Println("Timer Afer 5sec", time)
}

func TimerAfterFunc() {
	group := sync.WaitGroup{}
	group.Add(1)

	time.AfterFunc(5*time.Second, func() {
		fmt.Println("Timer After func 5sec", time.Now())
		group.Done()
	})
	fmt.Println("Timer After func", time.Now())

	group.Wait()
}

func main() {
	Timer()
	TimerAfter()
	TimerAfterFunc()
}
