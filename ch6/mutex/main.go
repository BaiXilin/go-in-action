package main

import (
	"fmt"
	"runtime"
	"sync"
)

var (
	counter int64

	wg sync.WaitGroup

	// mutex to define a zone in code
	mutex sync.Mutex
)

func main() {
	wg.Add(2)

	go incCounter(1)
	go incCounter(2)

	wg.Wait()

	fmt.Println("Final counter:", counter)

}

func incCounter(id int) {
	defer wg.Done()

	for count := 0; count < 2; count++ {
		// only one goroutine is allowed to enter this zone
		// at a time
		mutex.Lock()
		{
			value := counter
			runtime.Gosched()

			value++
			counter = value
		}
		mutex.Unlock()
	}
}
