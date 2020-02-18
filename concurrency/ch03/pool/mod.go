// sync.Pool example, page 59
package pool

import (
	"fmt"
	"sync"
)

var Pool = &sync.Pool{
	New: func() interface{} {
		fmt.Println("Creating new instance.")
		return struct{}{}
	},
}
