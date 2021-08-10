package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	"book/ch04"
)

var (
	addr = flag.String("addr", "", "listening address")
	port = flag.Int("port", 8008, "listening port")
)

func main() {
	flag.Parse()
	if *addr == "" {
		if len(os.Args) < 3 {
			usage()
		}
		nth, err := strconv.ParseInt(os.Args[1], 10, 0)
		if err != nil {
			fmt.Fprintf(os.Stderr, "rotate: %v\n", err)
			os.Exit(2)
		}
		a := make([]interface{}, 0, len(os.Args[2:]))
		for _, arg := range os.Args[2:] {
			a = append(a, arg)
		}
		ch04.Rotate(a, int(nth))
		fmt.Println(a)
		return
	}
	http.HandleFunc("/", handler)
	url := fmt.Sprintf("%s:%d", *addr, *port)
	log.Fatal(http.ListenAndServe(url, nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	args := strings.Split(r.URL.Path, "/")
	nth, err := strconv.ParseInt(args[1], 10, 0)
	if err != nil {
		log.Print(err)
		return
	}
	a := make([]interface{}, 0, len(args[1:]))
	for _, arg := range args[2:] {
		a = append(a, interface{}(arg))
	}
	ch04.Rotate(a, int(nth))
	fmt.Fprintln(w, a)
}

func usage() {
	flag.Usage()
	os.Exit(2)
}
