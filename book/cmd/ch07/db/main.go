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
		"shoes": 12.5,
		"socks": 8.99,
	}
	log.Fatal(http.ListenAndServe(*addr, db))
}

type dollar float32

func (d dollar) String() string {
	return fmt.Sprintf("$%.02f", d)
}

type database map[string]dollar

func (db database) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	switch req.URL.Path {
	case "/", "/list":
		db.list(w, req)
	case "/price":
		db.price(w, req)
	default:
		db.handle404(w, req)
	}
}

func (db database) list(w http.ResponseWriter, req *http.Request) {
	for item, price := range db {
		fmt.Fprintf(w, "%s: %s\n", item, price)
	}
}

func (db database) price(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	price, ok := db[item]
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "missing item: %q\n", item)
		return
	}
	fmt.Fprintf(w, "%s: %s\n", item, price)
}

func (db database) handle404(w http.ResponseWriter, req *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	fmt.Fprintf(w, "does not exist: %s\n", req.URL.Path)
}
