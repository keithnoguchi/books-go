package ch05

import (
	"fmt"
	"io"

	"golang.org/x/net/html"
)

func Title(data io.Reader) (*string, error) {
	doc, err := html.Parse(data)
	if err != nil {
		return nil, err
	}
	var forEachNode func(*html.Node) *string
	forEachNode = func(n *html.Node) *string {
		if n.Type == html.ElementNode && n.Data == "title" &&
			n.FirstChild != nil {
			return &n.FirstChild.Data
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			if title := forEachNode(c); title != nil {
				return title
			}
		}
		return nil
	}
	title := forEachNode(doc)
	if title == nil {
		return nil, fmt.Errorf("missing title")
	}
	return title, nil
}
