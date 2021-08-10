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
	port = flag.Int("port", 8012, "listening port")
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
	graph, err := ch04.CreateRuneGraph(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "graph: %v\n", err)
		return 1
	}
	fmt.Printf("FROM\tTOs\n")
	for from, e := range graph {
		fmt.Printf("%q\t[", from)
		var notFirst bool
		for to := range e {
			if notFirst {
				fmt.Print(",")
			}
			fmt.Printf("%q", to)
			notFirst = true
		}
		fmt.Printf("]\n")
	}
	return 0
}

func httpHandler(w http.ResponseWriter, r *http.Request) {
	graph, err := ch04.CreateRuneGraph(strings.NewReader(r.URL.Path))
	if err != nil {
		log.Printf("graph: %v", err)
		return
	}
	fmt.Fprintf(w, "FROM\tTOs\n")
	for from, e := range graph {
		fmt.Fprintf(w, "%q\t[", from)
		var notFirst bool
		for to := range e {
			if notFirst {
				fmt.Fprintf(w, ",")
			}
			fmt.Fprintf(w, "%q", to)
			notFirst = true
		}
		fmt.Fprintf(w, "]\n")
	}
}
