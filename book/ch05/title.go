package ch05

import (
	"fmt"
	"io"

	"golang.org/x/net/html"
)

func Title(in io.Reader) (title string, err error) {
	type bailout struct{}
	defer func() {
		switch p := recover(); p {
		case nil:
			// no panic
		case bailout{}:
			err = fmt.Errorf("multiple title")
		default:
			panic(p)
		}
	}()
	doc, err := html.Parse(in)
	if err != nil {
		return title, err
	}
	var forEachNode func(*html.Node)
	forEachNode = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "title" &&
			n.FirstChild != nil {
			if title != "" {
				panic(bailout{})
			}
			title = n.FirstChild.Data
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			forEachNode(c)
		}
	}
	forEachNode(doc)
	if title == "" {
		return title, fmt.Errorf("missing title")
	}
	return title, nil
}
