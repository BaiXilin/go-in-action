package pool

import (
	"errors"
	"io"
	"log"
	"sync"
)

// Pool manages resources that can be safely shared among goroutines
type Pool struct {
	m         sync.Mutex
	resources chan io.Closer            // where resources are stored, must impletement io.Closer interface
	factory   func() (io.Closer, error) // used to create  new resources, implemented by the caller
	closed    bool
}

var ErrPoolClosed = errors.New("Pool has  been closed.")

// New is the constructor of Pool, pass factory of resource in
func New(fn func() (io.Closer, error), size uint) (*Pool, error) {
	if size <= 0 {
		return nil, errors.New("Size value too small.")
	}

	return &Pool{
		factory:   fn,
		resources: make(chan io.Closer, size),
	}, nil
}

// Acquire gives resource from the pool
func (p *Pool) Acquire() (io.Closer, error) {
	select {
	// check if there are idle resources
	case r, ok := <-p.resources:
		log.Println("Acquire: Shared Resource")
		if !ok {
			return nil, ErrPoolClosed
		}
		return r, nil
	// there are  no idle resources, thus create a new one
	default:
		log.Println("Acquire: New Resource")
		return p.factory()
	}
}

// Release put a resource back to the pool after use
func (p *Pool) Release(r io.Closer) {
	// ensure this operation is safe with Close
	p.m.Lock()
	defer p.m.Unlock()

	// if the pool is closed, destroy this resource
	if p.closed {
		r.Close()
		return
	}

	select {
	// try to put this resource back into the pool
	case p.resources <- r:
		log.Println("Release: Back In Queue")
	// if the pool is full, destroy this resource
	default:
		log.Println("Release: Closing")
		r.Close()
	}
}

// Close will stop the pool from working, and close all existing resources
func (p *Pool) Close() {
	// ensure this operation is safe with Release
	p.m.Lock()
	defer p.m.Unlock()

	// if the pool is already closed, do nothing
	if p.closed {
		return
	}

	// close the pool
	p.closed = true

	// MUST close the channel before closing resources in it, or dead lock
	close(p.resources)

	// close resources
	for r := range p.resources {
		r.Close() // resource implemented io.Closer
	}
}
