// server4 outputs lissajous figures
package main

import (
	"log"
	"net/http"
	"os"

	server "book/ch01/server/v4"
)

const ADDR = "localhost:8000"

func main() {
	addr := ADDR
	if len(os.Args) > 1 {
		addr = os.Args[1]
	}
	http.HandleFunc("/", server.Handler)
	log.Fatal(http.ListenAndServe(addr, nil))
}
