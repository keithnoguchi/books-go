package server

import (
	"net/http"

	"book/ch01"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	ch01.Lissajous(w)
}
