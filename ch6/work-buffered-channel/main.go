// using buffered channel to do work with fixed number of goroutines
package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const (
	numberGoroutines = 4
	taskLoad         = 10
)

var wg sync.WaitGroup

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	// create a buffered channel to manage work
	tasks := make(chan string, taskLoad)

	// start Goroutines to do work
	wg.Add(numberGoroutines)
	for gr := 1; gr <= numberGoroutines; gr++ {
		go worker(tasks, gr)
	}

	// assign work to workers
	for post := 1; post <= taskLoad; post++ {
		tasks <- fmt.Sprintf("Task: %d", post)
	}

	// when all works are done, close the channel
	// so all goroutines can exit
	close(tasks)
	wg.Wait()
}

func worker(tasks chan string, worker int) {
	defer wg.Done()

	for {
		task, ok := <-tasks
		if !ok {
			// the channel is empty and closed
			fmt.Printf("Worker %d: Shutting down\n", worker)
			return
		}

		// start working
		fmt.Printf("Worker: %d : Started %s\n", worker, task)

		// wait a random long time to simulate work
		sleep := rand.Int63n(100)
		time.Sleep(time.Duration(sleep) * time.Millisecond)

		// show the work is done
		fmt.Printf("Worker: %d : Completed %s\n", worker, task)
	}
}
