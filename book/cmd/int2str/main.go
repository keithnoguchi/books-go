package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	"book/ch03"
)

var (
	addr = flag.String("addr", "", "listen address")
	port = flag.Int("port", 8006, "listen port")
)

func main() {
	flag.Parse()
	if *addr == "" {
		var ints []int
		for _, arg := range os.Args[1:] {
			v, err := strconv.ParseInt(arg, 10, 32)
			if err != nil {
				flag.Usage()
				os.Exit(1)
			}
			ints = append(ints, int(v))
		}
		fmt.Println(ch03.IntsToString(ints))
		return
	}
	url := fmt.Sprintf("%s:%d", *addr, *port)
	handler := func(w http.ResponseWriter, r *http.Request) {
		var ints []int
		for _, arg := range strings.Split(r.URL.Path, "/") {
			v, err := strconv.ParseInt(arg, 10, 0)
			if err != nil {
				log.Printf("invalid arg(%q)", arg)
				continue
			}
			ints = append(ints, int(v))
		}
		fmt.Fprintln(w, ch03.IntsToString(ints))
	}
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(url, nil))
}
