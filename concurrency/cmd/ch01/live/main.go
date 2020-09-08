package main

import (
	"sync"

	"concurrency/ch01/live"
)

func main() {
	var people sync.WaitGroup
	people.Add(2)
	go live.Walk(&people, "Alice")
	go live.Walk(&people, "Barbara")
	people.Wait()
}
