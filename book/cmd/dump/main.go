// dump dumps header info.
package main

import (
	"log"
	"net/http"

	"book/ch01"
)

func main() {
	log.Fatal(http.ListenAndServe("localhost:8002", &ch01.Dumper{}))
}
