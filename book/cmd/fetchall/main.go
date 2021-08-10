// fetchall fetches all the URL and returns the size of the body
package main

import (
	"fmt"
	"os"
	"sync"
	"time"

	"book/ch01"
)

func main() {
	start := time.Now()
	var wg sync.WaitGroup
	wg.Add(len(os.Args[1:]))
	ch := make(chan string)
	for _, url := range os.Args[1:] {
		go ch01.Fetch(&wg, ch, url)
	}
	for range os.Args[1:] {
		fmt.Print(<-ch)
	}
	wg.Wait()
	fmt.Printf("%.2fs elasped\n", time.Since(start).Seconds())
}
