package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

var (
	// notify the goroutines to stop working
	shutdown int64

	wg sync.WaitGroup
)

func main() {
	wg.Add(2)

	go doWork("A")
	go doWork("B")

	// give goroutines time to run
	time.Sleep(1 * time.Second)

	atomic.StoreInt64(&shutdown, 1)

	wg.Wait()
}

func doWork(name string) {
	defer wg.Done()

	for {
		fmt.Printf("Doing %s work\n", name)

		time.Sleep(250 * time.Millisecond)

		if atomic.LoadInt64(&shutdown) == 1 {
			fmt.Printf("Shutting %s down\n", name)
			break
		}
	}
}
