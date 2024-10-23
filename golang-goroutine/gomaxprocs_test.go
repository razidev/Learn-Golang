package goroutine

import (
	"fmt"
	"runtime"
	"sync"
	"testing"
	"time"
)

func TestGetGomaxprocs(t *testing.T) {
	group := sync.WaitGroup{}
	for i := 0; i < 100; i++ {
		group.Add(1)
		go func() {
			time.Sleep(3 * time.Second)
			group.Done()
		}()
	}

	fmt.Println("total cpu", runtime.NumCPU())
	fmt.Println("total thread", runtime.GOMAXPROCS(-1))
	fmt.Println("total goroutine", runtime.NumGoroutine())

	group.Wait()
}
