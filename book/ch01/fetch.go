package ch01

import (
	"fmt"
	"io"
	"net/http"
	"sync"
	"time"
)

func Fetch(wg *sync.WaitGroup, ch chan<- string, url string) {
	defer wg.Done()
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprintf("%s: %v\n", url, err)
		return
	}
	defer resp.Body.Close()
	n, err := io.Copy(io.Discard, resp.Body)
	if err != nil {
		ch <- fmt.Sprintf("%s: %v\n", url, err)
		return
	}
	ch <- fmt.Sprintf("%.2fs\t%d\t%s\n", time.Since(start).Seconds(), n, url)
}
