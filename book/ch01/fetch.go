// fetch URL contents
package ch01

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

func Fetch(url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err)
		return
	}
	defer resp.Body.Close()
	n, err := io.Copy(io.Discard, resp.Body)
	if err != nil {
		ch <- fmt.Sprint(err)
		return
	}
	secs := time.Since(start).Milliseconds()
	ch <- fmt.Sprintf("%4dms  %7dbytes  %s", secs, n, url)
}
