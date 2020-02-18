// Starvation example, page 16
package main

import (
	"sync"
	"time"

	"concurrency/ch01/greedy"
)

const runtime = 10 * time.Millisecond

func main() {
	var wg sync.WaitGroup
	wg.Add(2)
	go greedy.GreedyWorker(&wg, runtime)
	go greedy.PoliteWorker(&wg, runtime)
	wg.Wait()
}
