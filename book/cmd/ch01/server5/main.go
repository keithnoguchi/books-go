// server5 calls ch01.Lissajous function directly
package main

import (
	"log"
	"net/http"
	"os"

	"book/ch01"
)

const ADDR = "localhost:8000"

func main() {
	addr := ADDR
	if len(os.Args) > 1 {
		addr = os.Args[1]
	}
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		ch01.Lissajous(w)
	})
	log.Fatal(http.ListenAndServe(addr, nil))
}
