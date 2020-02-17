package main

import (
	"sync"

	"concurrency/ch01/live"
)

func main() {
	var peopleInHallway sync.WaitGroup
	peopleInHallway.Add(2)
	go live.Walk(&peopleInHallway, "Alice")
	go live.Walk(&peopleInHallway, "Barbara")
	peopleInHallway.Wait()
}
