package main

import (
	"fmt"
	"visibility_counter/counters"
)

func main() {
	counters := counters.New(10)

	fmt.Printf("Counter: %d of type %T\n", counters, counters)
}
