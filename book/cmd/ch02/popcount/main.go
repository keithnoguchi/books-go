// popcount counts the population count of command line arguments.
package main

import (
	"fmt"
	"os"
	"strconv"

	"book/ch02/popcount"
)

func main() {
	for _, arg := range os.Args[1:] {
		x, err := strconv.ParseUint(arg, 10, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "popcount: %v\n", err)
			os.Exit(1)
		}
		fmt.Printf("Population count of %q is %d\n",
			arg, popcount.PopCount(x))
	}
}
