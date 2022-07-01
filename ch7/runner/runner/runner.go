// runner manages the execution of tasks and their life cycle
package runner

import (
	"errors"
	"os"
	"os/signal"
	"time"
)

// runner executes tasks with given timeout
// tasks are finished when os send interrupt signal
type Runner struct {
	// signal from os
	interrupt chan os.Signal

	// complete used to report tasks are finished
	complete chan error

	// timeout used to report the tasks are over time
	timeout <-chan time.Time

	// a groups of tasks
	tasks []func(int)
}

// ErrTimeout returns when tasks are timeout
var ErrTimeout = errors.New("received timeout")

// ErrInterrupt returns when received interrupt signal from os
var ErrInterrupt = errors.New("received interrupt")

// runner constructor
func New(d time.Duration) *Runner {
	return &Runner{
		interrupt: make(chan os.Signal, 1),
		complete:  make(chan error),
		timeout:   time.After(d),
	}
}

// add tasks to runner
func (r *Runner) Add(tasks ...func(int)) { // ... for variadic functions
	r.tasks = append(r.tasks, tasks...)
}

// run executes tasks in a runner
func (r *Runner) run() error {
	for id, task := range r.tasks {
		if r.gotInterrupt() {
			return ErrInterrupt
		}
		task(id)
	}
	return nil
}

// !!! classic usage of select and default
// gotInterrupt verifies if receive a interrupt signal
func (r *Runner) gotInterrupt() bool {
	select {
	case <-r.interrupt:
		// stop receiving furthur signals
		signal.Stop(r.interrupt)
		return true
	default: // default makes the select non-blocking
		return false
	}
}

// execute all tasks and monitor events on channels
func (r *Runner) Start() error {
	// we want to receive all interrupt signal
	signal.Notify(r.interrupt, os.Interrupt)

	// run goroutines to execute tasks
	go func() {
		r.complete <- r.run()
	}()

	select {
	case err := <-r.complete:
		return err
	case <-r.timeout:
		return ErrTimeout
	}
}
