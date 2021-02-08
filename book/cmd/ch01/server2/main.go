// server2 demonstrates the sync.Mutex usage
package main

import (
	"log"
	"net/http"
	"os"

	server "book/ch01/server2"
)

const ADDR = "localhost:8000"

func main() {
	addr := ADDR
	if len(os.Args) > 1 {
		addr = os.Args[1]
	}
	http.HandleFunc("/", server.Handler)
	http.HandleFunc("/metrics", server.Counter)
	log.Fatal(http.ListenAndServe(addr, nil))
}
