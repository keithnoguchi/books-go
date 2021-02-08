// echo4 prints its command line arguments.
package main

import (
	"flag"
	"fmt"
	"strings"
)

func main() {
	var n = flag.Bool("n", false, "omit trailing newline")
	var sep = flag.String("s", " ", "separator")
	flag.Parse()
	fmt.Print(strings.Join(flag.Args(), *sep))
	if len(flag.Args()) > 0 && !*n {
		fmt.Println()
	}
}
