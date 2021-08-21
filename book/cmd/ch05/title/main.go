package main

import (
	"fmt"
	"net/http"
	"os"

	"book/ch05"
)

func main() {
	title := func(url string) (*string, error) {
		resp, err := http.Get(url)
		if err != nil {
			return nil, err
		}
		defer resp.Body.Close()
		return ch05.Title(resp.Body)
	}
	for _, url := range os.Args[1:] {
		title, err := title(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "title: %v\n", err)
			continue
		}
		fmt.Println(*title)
	}
}

