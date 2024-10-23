package goroutine

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func RunAsynchronous(group *sync.WaitGroup, number int) {
	defer group.Done()

	group.Add(1)

	fmt.Println("Hello", number)
	time.Sleep(1 * time.Second)
}

func TestWaitGroup(t *testing.T) {
	group := &sync.WaitGroup{}

	for i := 0; i < 100; i++ {
		fmt.Println(i)
		go RunAsynchronous(group, i)
	}

	group.Wait()
	fmt.Println("done")
}
