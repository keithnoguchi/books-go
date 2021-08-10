package ch01

import (
	"fmt"

	"net/http"
	"sync/atomic"
)

type Counter struct {
	count uint32
}

func (c *Counter) Count(w http.ResponseWriter, req *http.Request) {
	atomic.AddUint32(&c.count, 1)
}

func (c *Counter) GetCount(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "count = %d!\n", atomic.LoadUint32(&c.count))
}
