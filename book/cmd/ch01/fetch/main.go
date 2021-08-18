package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, "fetch <URL>")
		os.Exit(1)
	}
	resp, err := http.Get(os.Args[1])
	if err != nil {
		fmt.Fprintln(os.Stdout, err)
		os.Exit(1)
	}
	b, err := io.ReadAll(resp.Body)
	defer resp.Body.Close()
	fmt.Printf("%s", b)
}
