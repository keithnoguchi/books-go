// echo4: just fmt.Println()
package main

import (
	"fmt"
	"os"
)

func main() {
	// It prints as [arg1 arg2 arg3] but it's quick.
	fmt.Println(os.Args[1:])
}
