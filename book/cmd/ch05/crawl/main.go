package main

import (
	"fmt"
	"os"

	"book/ch05"
)

func main() {
	ch05.BreathFirst(crawl, os.Args[1:])
}

func crawl(url string) []string {
	fmt.Println(url)
	urls, err := ch05.Extract(url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "crawl(%q): %v\n", url, err)
		return nil
	}
	return urls
}
