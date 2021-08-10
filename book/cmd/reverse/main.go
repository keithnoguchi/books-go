package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"book/ch04"
)

var (
	addr = flag.String("addr", "", "listening address")
	port = flag.Int("port", 8007, "listening port")
)

func main() {
	flag.Parse()
	if *addr == "" {
		var a []interface{}
		for _, arg := range os.Args[1:] {
			a = append(a, interface{}(arg))
		}
		ch04.Reverse(a)
		fmt.Println(a)
		return
	}
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		var a []interface{}
		for _, arg := range strings.Split(r.URL.Path, "/") {
			a = append(a, interface{}(arg))
		}
		ch04.Reverse(a)
		fmt.Fprintln(w, a)
	})
	url := fmt.Sprintf("%s:%d", *addr, *port)
	log.Fatal(http.ListenAndServe(url, nil))
}
