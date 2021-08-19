package ch05

import (
	"fmt"
	"io"

	"golang.org/x/net/html"
)

func Outline(w io.Writer, stack []string, n *html.Node) {
	if n.Type == html.ElementNode {
		stack = append(stack, n.Data)
		fmt.Fprintln(w, stack)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		Outline(w, stack, c)
	}
}
