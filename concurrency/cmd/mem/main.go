// Memory pool example
package main

import (
	"fmt"
	"sync"
	"time"

	"concurrency/ch03/pool"
)

func init() {
	// Seed the pool with 4KB
	for i := 0; i < 4; i++ {
		pool.CalcPool.Put(pool.CalcPool.New())
	}
}

func main() {
	const numWorkers = 1024 * 1024
	var wg sync.WaitGroup
	wg.Add(numWorkers)
	for i := numWorkers; i > 0; i-- {
		go func() {
			defer wg.Done()
			mem := pool.CalcPool.Get().(*[]byte)
			defer pool.CalcPool.Put(mem)
			// Something important but quick to be done...
			time.Sleep(1 * time.Nanosecond)
		}()
	}
	wg.Wait()
	fmt.Printf("%d calculators were created.\n", pool.NumCalcsCreated)
}
