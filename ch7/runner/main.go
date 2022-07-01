package main

import (
	"log"
	"os"
	"time"

	"example.com/baixilin/runner-main/runner"
)

const timeout = 5 * time.Second

func main() {
	log.Println("Starting work")

	// assign timeout
	r := runner.New(timeout)

	// add tasks to execute
	r.Add(createTask(), createTask(), createTask())

	// execute tasks and handle the results
	if err := r.Start(); err != nil {
		switch err {
		case runner.ErrTimeout:
			log.Println("Terminating due to timeout")
			os.Exit(1)
		case runner.ErrInterrupt:
			log.Println("Terminating due to interrupt")
			os.Exit(2)
		}
	}
	log.Println("Process ended.")
}

func createTask() func(int) {
	return func(id int) {
		log.Printf("Processor  - Task #%d.", id)
		time.Sleep(time.Duration(id) * time.Second)
	}
}
