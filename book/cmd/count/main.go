// count: counts the requests
package main

import (
	"log"
	"net/http"

	"book/ch01"
)

func main() {
	counter := &ch01.Counter{}
	http.HandleFunc("/", counter.Count)
	http.HandleFunc("/count", counter.GetCount)
	log.Fatal(http.ListenAndServe("localhost:8001", nil))
}
