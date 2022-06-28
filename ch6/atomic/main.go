package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

var (
	// Every Goroutine will try to add to counter
	counter int64

	wg sync.WaitGroup
)

func main() {
	wg.Add(2)

	go incCounter(1)
	go incCounter(2)

	wg.Wait()

	fmt.Println("Final Counter:", counter)
}

func incCounter(id int) {
	defer wg.Done()

	for count := 0; count < 2; count++ {
		// safely add to counter
		atomic.AddInt64(&counter, 1)

		runtime.Gosched()
	}
}
