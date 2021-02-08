package server

import (
	"fmt"
	"net/http"
)

func Handler(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "URL.Path = %q\n", req.URL.Path)
}
