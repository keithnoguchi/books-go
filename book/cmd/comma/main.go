package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"book/ch03/comma"
)

var (
	addr = flag.String("addr", "", "listen address")
	port = flag.Int("port", 8005, "listen port")
)

func main() {
	flag.Parse()
	if *addr == "" {
		if len(os.Args) > 1 {
			fmt.Println(comma.Comma(os.Args[1]))
		} else {
			fmt.Println(comma.Comma(os.Args[0]))
		}
		return
	}
	url := fmt.Sprintf("%s:%d", *addr, *port)
	handler := func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, comma.Comma(r.URL.Path))
	}
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(url, nil))
}
