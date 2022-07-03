package main

import (
	"io"
	"log"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"

	"example.com/baixilin/pool-main/pool"
)

const (
	maxGoroutines  = 25
	pooledResource = 2
)

// mock resource
type dbConnection struct {
	ID int32
}

func (dbConn *dbConnection) Close() error {
	log.Println("Close: Connection", dbConn.ID)
	return nil
}

var idCounter int32

func createDbConn() (io.Closer, error) {
	id := atomic.AddInt32(&idCounter, 1)
	log.Println("Create: Connection", idCounter)
	return &dbConnection{id}, nil
}

func main() {
	var wg sync.WaitGroup
	wg.Add(maxGoroutines)

	// create resource pool
	p, err := pool.New(createDbConn, pooledResource)
	if err != nil {
		log.Println(err)
	}

	// use resources from pool to do query search
	for query := 0; query < maxGoroutines; query++ {
		go func(q int) {
			performQueries(q, p)
			wg.Done()
		}(query)
	}

	wg.Wait()
	log.Println("Shutdown")
	p.Close()
}

func performQueries(query int, p *pool.Pool) {
	// get a resource from the pool
	conn, err := p.Acquire()
	if err != nil {
		log.Println(err)
		return
	}

	// after done using, put back to pool
	defer p.Release(conn)

	// mock using the resource
	time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
	log.Printf("QID[%d] CID[%d]\n", query, conn.(*dbConnection).ID)
}
