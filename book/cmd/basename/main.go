package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	basename "book/ch03/basename/v2"
)

var (
	addr = flag.String("addr", "", "listening address")
	port = flag.Int("port", 8003, "listening port")
)

func main() {
	flag.Parse()
	if *addr != "" {
		url := fmt.Sprintf("%s:%d", *addr, *port)
		handler := func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintln(w, basename.Basename(r.URL.Path))
		}
		http.HandleFunc("/", handler)
		log.Fatal(http.ListenAndServe(url, nil))
	}
	var path *string
	if len(flag.Args()) > 0 {
		path = &flag.Args()[0]
	} else {
		path = &os.Args[0]
	}
	fmt.Println(basename.Basename(*path))
}
