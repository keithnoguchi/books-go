package server

import (
	"fmt"
	"net/http"
	"sync"
)

func Counter(w http.ResponseWriter, req *http.Request) {
	counterLock.Lock()
	defer counterLock.Unlock()
	fmt.Fprintf(w, "Counter = %d\n", counter)
}

var counterLock sync.Mutex
var counter uint64

func inc() uint64 {
	counterLock.Lock()
	defer counterLock.Unlock()
	counter++
	return counter
}
