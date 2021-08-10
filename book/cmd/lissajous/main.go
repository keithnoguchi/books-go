// lissajous: Let's draw a lissajous image
package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"book/ch01"
)

var (
	address = flag.String("address", "", "listening address")
	port    = flag.Int("port", 8000, "listening port")
)

func main() {
	flag.Parse()
	if *address != "" {
		addr := fmt.Sprintf("%s:%d", *address, *port)
		http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
			ch01.Lissajous(w)
		})
		log.Fatal(http.ListenAndServe(addr, nil))
	}
	ch01.Lissajous(os.Stdout)
}
