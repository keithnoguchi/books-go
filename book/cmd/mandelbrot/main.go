package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"book/ch03/mandelbrot"
)

var (
	addr = flag.String("address", "", "listening address")
	port = flag.Int("port", 8002, "listening port")
)

func main() {
	flag.Parse()
	if *addr == "" {
		mandelbrot.Draw(os.Stdout)
		return
	}
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		mandelbrot.Draw(w)
	})
	url := fmt.Sprintf("%s:%d", *addr, *port)
	log.Fatal(http.ListenAndServe(url, nil))
}
