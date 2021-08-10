// read lines from stdin and print out the duplicate lines to stdout
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)

	// read lines from scanner
	for input.Scan() {
		counts[input.Text()]++
	}

	// print the result.
	for line, n := range counts {
		if n < 2 {
			continue
		}
		fmt.Printf("%d\t%q\n", n, line)
	}
}
