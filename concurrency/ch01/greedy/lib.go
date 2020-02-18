// Starvation example, page 16
package greedy

import (
	"fmt"
	"sync"
	"time"
)

var sharedLock sync.Mutex

func GreedyWorker(wg *sync.WaitGroup, runtime time.Duration) {
	defer wg.Done()
	var count int
	for begin := time.Now(); time.Since(begin) <= runtime; {
		sharedLock.Lock()
		time.Sleep(3 * time.Nanosecond)
		sharedLock.Unlock()
		count++
	}
	fmt.Printf(
		"Greedy worker was able to execute %v work loops\n",
		count,
	)
}

func PoliteWorker(wg *sync.WaitGroup, runtime time.Duration) {
	defer wg.Done()
	var count int
	for begin := time.Now(); time.Since(begin) <= runtime; {
		sharedLock.Lock()
		time.Sleep(1 * time.Nanosecond)
		sharedLock.Unlock()

		sharedLock.Lock()
		time.Sleep(1 * time.Nanosecond)
		sharedLock.Unlock()

		sharedLock.Lock()
		time.Sleep(1 * time.Nanosecond)
		sharedLock.Unlock()

		count++
	}
	fmt.Printf(
		"Polite worker was able to execute %v work loops.\n",
		count,
	)
}
