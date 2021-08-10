// SPDX-License-Identifier: GPL-2.0
package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"book/ch04"
)

var (
	host = flag.String("host", "", "listening host, e.g. localhost")
	port = flag.Int("port", 8011, "listening port")
)

func main() {
	flag.Parse()
	if *host == "" {
		os.Exit(cmdHandler())
	}
	addr := fmt.Sprintf("%s:%d", *host, *port)
	http.HandleFunc("/", httpHandler)
	log.Fatal(http.ListenAndServe(addr, nil))
}

func cmdHandler() int {
	c := ch04.NewCharCounter()
	if err := c.Count(os.Stdin); err != nil {
		fmt.Fprintf(os.Stderr, "count: %v\n", err)
		return 1
	}
	fmt.Printf("rune\tcount\n")
	for r, n := range c.Counts {
		fmt.Printf("%q\t%d\n", r, n)
	}
	fmt.Printf("\nlen\tcount\n")
	for n, c := range c.UTFLen() {
		if n > 0 {
			fmt.Printf("%d\t%d\n", n, c)
		}
	}
	if invalid := c.InvalidCount(); invalid > 0 {
		fmt.Printf("\n%d invalid UTF-8 characters\n", invalid)
	}
	return 0
}

func httpHandler(w http.ResponseWriter, r *http.Request) {
	c := ch04.NewCharCounter()
	if err := c.Count(strings.NewReader(r.URL.Path)); err != nil {
		log.Printf("charcount: %v\n", err)
		return
	}
	fmt.Fprintf(w, "rune\tcount\n")
	for r, n := range c.Counts {
		fmt.Fprintf(w, "%q\t%d\n", r, n)
	}
	fmt.Fprintf(w, "\nlen\tcount\n")
	for n, c := range c.UTFLen() {
		fmt.Fprintf(w, "%d\t%d\n", n, c)
	}
	if invalid := c.InvalidCount(); invalid != 0 {
		fmt.Fprintf(w, "\n%d invalid UTF-8 characters\n", invalid)
	}
}
