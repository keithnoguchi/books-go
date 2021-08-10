// path: simple path web server
package main

import (
	"log"
	"net/http"

	"book/ch01"
)

func main() {
	http.HandleFunc("/", ch01.PrintPath)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}
