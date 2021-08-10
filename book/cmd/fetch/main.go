// fetch the data from url
package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

const progname string = "fetch1"

func main() {
	for _, url := range os.Args[1:] {
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s: %v\n", progname, err)
			os.Exit(1)
		}
		defer resp.Body.Close()
		b, err := io.ReadAll(resp.Body)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s: %v\n", progname, err)
			os.Exit(1)
		}
		fmt.Print(string(b))
	}
}
