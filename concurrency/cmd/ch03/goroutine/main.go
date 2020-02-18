// Goroutine memory usage, page 43
package main

import (
	"fmt"
	"sync"

	"concurrency/ch03/goroutine"
)

const numGoroutines = 1e4

func main() {
	var wg sync.WaitGroup
	wg.Add(numGoroutines)
	before := goroutine.MemConsumed()
	for i := numGoroutines; i > 0; i-- {
		goroutine.Noop(&wg)
	}
	wg.Wait()
	after := goroutine.MemConsumed()
	fmt.Printf("%.3fBytes\n", float64(after-before)/numGoroutines)
}
