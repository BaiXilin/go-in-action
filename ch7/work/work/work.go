package work

import "sync"

// specify the interface type of worker to use the pool
type Worker interface {
	Task()
}

// the goroutine pool that can assign to Worker tasks
type Pool struct {
	work chan Worker // why unbuffered channel?
	wg   sync.WaitGroup
}

// New creates a new Pool
func New(maxGoroutines int) *Pool {
	p := Pool{
		work: make(chan Worker),
	}

	p.wg.Add(maxGoroutines)
	for i := 0; i < maxGoroutines; i++ {
		go func() {
			for w := range p.work {
				w.Task()
			}
			p.wg.Done()
		}()
	}

	return &p
}

// Run submits work to pool
func (p *Pool) Run(w Worker) {
	p.work <- w
}

// Shutdown waits for all goroutines to stop
func (p *Pool) Shutdown() {
	close(p.work)
	p.wg.Wait()
}
