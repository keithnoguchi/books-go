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
	port = flag.Int("port", 8009, "listening port")
)

func main() {
	flag.Parse()
	if *addr != "" {
		url := fmt.Sprintf("%s:%d", *addr, *port)
		http.HandleFunc("/", httpHandler)
		log.Fatal(http.ListenAndServe(url, nil))
	}
	cmdHandler()
}

func httpHandler(w http.ResponseWriter, r *http.Request) {
	args := strings.Split(r.URL.Path, "/")
	nth, err := strconv.ParseInt(args[1], 10, 0)
	if err != nil {
		log.Println(err)
		return
	}
	a := make([]interface{}, 0, len(args[2:]))
	for _, arg := range args[2:] {
		a = append(a, arg)
	}
	a = ch04.Remove(a, int(nth))
	fmt.Fprintln(w, a)
}

func cmdHandler() {
	if len(os.Args) < 3 {
		flag.Usage()
		os.Exit(1)
	}
	nth, err := strconv.ParseInt(os.Args[1], 10, 0)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	var a []interface{}
	for _, v := range os.Args[2:] {
		a = append(a, v)
	}
	b := ch04.Remove(a, int(nth))
	fmt.Printf("%v\n", b)
	fmt.Printf("&a=%p,a=%p\n", &a, a)
	fmt.Printf("&b=%p,b=%p\n", &b, b)
	os.Exit(0)
}
