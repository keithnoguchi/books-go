// fetchall: fetch the contents of urls concurrently
package main

import (
	"fmt"
	"os"
	"time"

	"book/ch01"
)

func main() {
	start := time.Now()
	ch := make(chan string)
	for _, url := range os.Args[1:] {
		go ch01.Fetch(url, ch)
	}
	for range os.Args[1:] {
		fmt.Println(<-ch)
	}
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}
