// tostring: Change the integers to strings
package main

import (
	"fmt"
	"os"
	"strconv"

	"book/ch03"
)

func main() {
	if len(os.Args) < 1 {
		fmt.Fprintf(os.Stderr, "usage: tostring <integers...>\n")
		os.Exit(1)
	}
	var ints []int
	for _, v := range os.Args[1:] {
		n, err := strconv.Atoi(v)
		if err != nil {
			fmt.Fprintf(os.Stderr, "invalid argument: %q\n", v)
			os.Exit(1)
		}
		ints = append(ints, n)
	}
	fmt.Println(ch03.IntsToString(ints))
}
