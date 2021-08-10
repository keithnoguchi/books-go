// SPDX-License-Identifier: GPL-2.0
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	files := os.Args[1:]

	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for _, file := range files {
			f, err := os.Open(file)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts)
			f.Close()
		}
	}

	for line, n := range counts {
		if n < 2 {
			continue
		}
		fmt.Printf("%d\t%q\n", n, line)
	}
}

func countLines(f *os.File, counts map[string]int) {
	line := bufio.NewScanner(f)
	for line.Scan() {
		counts[line.Text()]++
	}
}
