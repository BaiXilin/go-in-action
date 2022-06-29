// using an unbuffered channel to simulate tennis match
// between two goroutines
package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var wg sync.WaitGroup

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	// create an unbuffered channel using make
	court := make(chan int)

	wg.Add(2)

	go player("Alice", court)
	go player("Bob", court)

	// match starts!
	court <- 1

	// wait until game over
	wg.Wait()
}

func player(name string, court chan int) {
	defer wg.Done()

	for {
		// wait for in-coming ball
		ball, ok := <-court
		if !ok {
			// if channel is closed, this player wins
			fmt.Printf("Player %s won\n", name)
			return
		}

		// random number generator to determine if this player loses
		n := rand.Intn(100)
		if n%13 == 0 {
			fmt.Printf("Player %s missed\n", name)
			close(court)
			return
		}

		// if not missed or won, display current status
		fmt.Printf("Player %s hit %d\n", name, ball)
		ball++

		// hit the ball to oppoent
		court <- ball
	}
}
