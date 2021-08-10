// surface: computes an SVG rendering of a 3-D surface function
package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"book/ch03/surface"
)

var (
	addr = flag.String("address", "", "listening address")
	port = flag.Int("port", 8001, "listening port")
)

func main() {
	flag.Parse()
	if *addr == "" {
		surface.Draw(os.Stdout)
		return
	}
	url := fmt.Sprintf("%s:%d", *addr, *port)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "image/svg+xml")
		surface.Draw(w)
	})
	log.Fatal(http.ListenAndServe(url, nil))
}
