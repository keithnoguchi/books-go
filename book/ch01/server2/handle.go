package server

import (
	"fmt"
	"net/http"
)

func Handler(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "%d> req.URL.Path = %q\n", inc(), req.URL.Path)
}
