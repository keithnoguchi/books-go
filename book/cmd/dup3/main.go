// dup3: read files into memory with os.ReadFile() and parse it.
package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	counts := make(map[string]int)
	for _, fname := range os.Args[1:] {
		data, err := os.ReadFile(fname)
		if err != nil {
			fmt.Fprintf(os.Stderr, "dup3: %v\n", err)
			continue
		}
		for _, line := range strings.Split(string(data), "\n") {
			counts[line]++
		}
	}
	for line, n := range counts {
		if n < 2 {
			continue
		}
		fmt.Printf("%d\t%q\n", n, line)
	}
}
