package server

import (
	"fmt"
	"log"
	"net/http"
)

func Handler(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "%s %s %s\n", req.Method, req.URL, req.Proto)
	for k, v := range req.Header {
		fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
	}
	fmt.Fprintf(w, "Host = %q\n", req.Host)
	fmt.Fprintf(w, "RemoteAddr = %q\n", req.RemoteAddr)
	if err := req.ParseForm(); err != nil {
		log.Print(err)
	}
	for k, v := range req.Form {
		fmt.Fprintf(w, "Form[%q] = %q\n", k, v)
	}
}
