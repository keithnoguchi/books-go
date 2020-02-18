// sync.Pool examples, page 59
package pool

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var Pool = &sync.Pool{
	New: func() interface{} {
		fmt.Println("Creating new instance.")
		return struct{}{}
	},
}

var NumCalcsCreated uint64

var CalcPool = &sync.Pool{
	New: func() interface{} {
		atomic.AddUint64(&NumCalcsCreated, 1)
		mem := make([]byte, 1024)
		return &mem
	},
}
