package live

import (
	"bytes"
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

var cadence = sync.NewCond(&sync.Mutex{})
var left, right int32

func init() {
	go func() {
		for range time.Tick(1 * time.Millisecond) {
			cadence.Broadcast()
		}
	}()
}

func Walk(wg *sync.WaitGroup, name string) {
	var out bytes.Buffer
	defer func() { fmt.Println(out.String()) }()
	defer wg.Done()
	fmt.Fprintf(&out, "%s is trying to scoot:", name)
	for i := 0; i < 5; i++ {
		if tryLeft(&out) || tryRight(&out) {
			return
		}
	}
	fmt.Fprintf(&out, "\n%s tosses her hands up in exasperation!", name)
}

func tryLeft(out *bytes.Buffer) bool {
	return tryDir("left", &left, out)
}

func tryRight(out *bytes.Buffer) bool {
	return tryDir("right", &right, out)
}

func tryDir(dirName string, dir *int32, out *bytes.Buffer) bool {
	fmt.Fprintf(out, " %s", dirName)
	atomic.AddInt32(dir, 1)
	takeStep()
	if atomic.LoadInt32(dir) == 1 {
		fmt.Fprint(out, ". Success!")
		return true
	}
	takeStep()
	atomic.AddInt32(dir, -1)
	return false
}

func takeStep() {
	cadence.L.Lock()
	cadence.Wait()
	cadence.L.Unlock()
}
