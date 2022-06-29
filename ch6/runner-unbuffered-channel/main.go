// using an unbuffered channel to simulate 4 runners
package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	baton := make(chan int)

	wg.Add(1)

	go Runner(baton)

	baton <- 1

	wg.Wait()

}

func Runner(baton chan int) {
	var newRunner int

	runner := <-baton

	fmt.Printf("Runner %d running with baton\n", runner)

	// create the next runner
	if runner != 4 {
		newRunner = runner + 1
		fmt.Printf("Runner %d to the line\n", newRunner)
		go Runner(baton)
	}

	// running on the field
	time.Sleep(1000 * time.Millisecond)

	// if the match over?
	if runner == 4 {
		fmt.Printf("Runner %d finished. Race over!\n", runner)
		wg.Done()
		return
	}

	// pass the baton to the next runner
	fmt.Printf("Runner %d pass baton to runner %d\n", runner, newRunner)
	baton <- newRunner
}
