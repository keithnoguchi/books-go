package ch01

import (
	"fmt"
	"net/http"
)

func PrintPath(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "URL = %q\n", r.URL)
}
