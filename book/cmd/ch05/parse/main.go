package main

import (
	"fmt"
	"net/http"
	"os"

	"golang.org/x/net/html"

	"book/ch05"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, "usage: parse URL\n")
		os.Exit(1)
	}
	resp, err := http.Get(os.Args[1])
	if err != nil {
		fmt.Fprintf(os.Stderr, "parse: %v\n", err)
		os.Exit(1)
	}
	defer resp.Body.Close()
	top, err := html.Parse(resp.Body)
	if err != nil {
		fmt.Fprintf(os.Stderr, "parse: %v\n", err)
		os.Exit(1)
	}
	ch05.ForEachNode(top, pre, post)
}

var depth uint

func pre(n *html.Node) {
	if n.Type != html.ElementNode {
		return
	}
	fmt.Printf("%*s<%s>\n", depth, "", n.Data)
	depth++
}

func post(n *html.Node) {
	if n.Type != html.ElementNode {
		return
	}
	depth--
	fmt.Printf("%*s</%s>\n", depth, "", n.Data)
}
