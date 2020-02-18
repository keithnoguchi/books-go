// Goroutine memory usage example, page 43
package goroutine

import (
	"runtime"
	"sync"
	"time"
)

var c <-chan interface{}

// Spawns the leaking goroutine.
func Noop(wg *sync.WaitGroup) {
	go func() {
		wg.Done()
		time.Sleep(1 * time.Second)
		// blocks here.
		<-c
	}()
}

// MemConsumed returns the current system memory consumed.
func MemConsumed() uint64 {
	runtime.GC()
	var s runtime.MemStats
	runtime.ReadMemStats(&s)
	return s.Sys
}
