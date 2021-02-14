package ch04

import (
	"bufio"
	"fmt"
	"io"
)

func Dedup(in io.Reader, out io.Writer) {
	seen := make(map[string]bool)
	input := bufio.NewScanner(in)
	for input.Scan() {
		line := input.Text()
		if !seen[line] {
			seen[line] = true
			fmt.Fprintln(out, line)
		}
	}
}
