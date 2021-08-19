package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
	"book/ch05"
)

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "outline: %v\n", err)
		os.Exit(1)
	}
	ch05.Outline(os.Stdout, nil, doc)
}
