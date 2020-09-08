// Deadlock, page 10
package main

import (
	"fmt"
	"sync"
	"time"
)

type value struct {
	value int
	lock  sync.Mutex
}

func main() {
	var wg sync.WaitGroup
	printSum := func(a, b *value) {
		defer wg.Done()
		a.lock.Lock()
		defer a.lock.Unlock()

		time.Sleep(2 * time.Second)
		b.lock.Lock()
		defer b.lock.Unlock()
		fmt.Printf("sum=%v\n", a.value+b.value)
	}

	var a, b value
	wg.Add(2)
	go printSum(&a, &b)
	go printSum(&b, &a)
	wg.Wait()
}
