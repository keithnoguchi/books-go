// SPDX-License-Identifier: GPL-2.0
package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
)

var addr = flag.String("addr", ":8080", "listening address")

func main() {
	flag.Parse()
	db := database{
		"shoes":    9.44,
		"t-shirts": 15.99,
	}
	log.Fatal(http.ListenAndServe(*addr, db))
}

type database map[string]dollar

func (db database) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	switch req.URL.Path {
	case "/", "/list":
		for item, price := range db {
			fmt.Fprintf(w, "%s: %s\n", item, price)
		}
	case "/price":
		item := req.URL.Query().Get("item")
		price, ok := db[item]
		if !ok {
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprintf(w, "item not found: %q\n", item)
			return
		}
		fmt.Fprintf(w, "%s: %s\n", item, price)
	default:
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "path not found: %s\n", req.URL.Path)
	}
}

type dollar float32

func (d dollar) String() string {
	return fmt.Sprintf("$%.02f", d)
}
