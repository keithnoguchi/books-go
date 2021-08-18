package main

import (
	"fmt"
	"os"

	"book/ch05"
	"golang.org/x/net/html"
)

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlink: %v\n", err)
		os.Exit(1)
	}
	for _, link := range ch05.Visit(nil, doc) {
		fmt.Println(link)
	}
}
