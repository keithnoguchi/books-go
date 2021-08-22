package ch05

import (
	"fmt"
	"net/http"

	"golang.org/x/net/html"
)

func Extract(url string) ([]string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("getting %s: %s", url, resp.Status)
	}
	doc, err := html.Parse(resp.Body)
	if err != nil {
		return nil, err
	}
	var list []string
	visitNode := func(n *html.Node) {
		if n.Type != html.ElementNode || n.Data != "a" {
			return
		}
		for _, attr := range n.Attr {
			if attr.Key != "href" {
				continue
			}
			link, err := resp.Request.URL.Parse(attr.Val)
			if err != nil {
				continue // ignore error
			}
			list = append(list, link.String())
		}
	}
	forEachNode(doc, visitNode)
	return list, nil
}

func forEachNode(n *html.Node, pre func(*html.Node)) {
	if pre != nil {
		pre(n)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, pre)
	}
}
