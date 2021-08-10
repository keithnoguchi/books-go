// SPDX-License-Identifier: GPL-2.0
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
	host = flag.String("host", "", "listening host")
	port = flag.Int("port", 8012, "listening port")
)

func main() {
	flag.Parse()
	if *host == "" {
		os.Exit(cmdHandler())
	}
	http.HandleFunc("/", httpHandler)
	addr := fmt.Sprintf("%s:%d", *host, *port)
	log.Fatal(http.ListenAndServe(addr, nil))
}

func httpHandler(w http.ResponseWriter, r *http.Request) {
	args := strings.Split(r.URL.Path, "/")
	ch04.Sort(args)
	fmt.Fprintln(w, args)
}

func cmdHandler() int {
	args := os.Args[1:]
	ch04.Sort(args)
	fmt.Println(args)
	return 0
}
