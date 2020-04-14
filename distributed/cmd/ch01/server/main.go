package main

import (
	"log"

	"distributed/ch01"
)

func main() {
	srv := ch01.NewHTTPServer(":8080")
	log.Fatal(srv.ListenAndServe())
}
