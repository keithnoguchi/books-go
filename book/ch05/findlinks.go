package ch05

import (
	"log"
	"net/http"
	"os"

	"golang.org/x/net/html"
)

func FindLinks(url string) ([]string, error) {
	log.SetPrefix("ch05.FindLinks: ")
	log.SetOutput(os.Stderr)

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	doc, err := html.Parse(resp.Body)
	if err != nil {
		return nil, err
	}
	links := visit(nil, doc)
	return links, nil
}

func visit(links []string, n *html.Node) []string {
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, attr := range n.Attr {
			if attr.Key == "href" {
				if len(attr.Val) == 0 {
					log.Println("zero length url")
					continue
				}
				links = append(links, attr.Val)
			}
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		links = visit(links, c)
	}
	return links
}
