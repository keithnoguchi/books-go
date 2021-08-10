// pc: counting set bits in 64 bit variable
package main

import (
	"fmt"
	"os"
	"strconv"

	"book/ch02/popcount"
)

func main() {
	for _, arg := range os.Args[1:] {
		val, err := strconv.ParseUint(arg, 10, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "pc: %v\n", err)
			os.Exit(1)
		}
		fmt.Printf(
			"%d number of bits in %q\n",
			popcount.Uint64(val),
			arg,
		)
	}
}
