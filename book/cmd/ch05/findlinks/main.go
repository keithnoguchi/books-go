// SPDX-License-Identifier: GPL-2.0
package main

import (
	"fmt"
	"os"

	"book/ch05"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, "usage: findlinks URL")
		os.Exit(1)
	}
	for _, url := range os.Args[1:] {
		links, err := ch05.FindLinks(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "findlinks(%s): %v\n", url, err)
			continue
		}
		for _, link := range links {
			fmt.Println(link)
		}
	}
}
