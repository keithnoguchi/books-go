// echo prints out the command line arguments
package main

import (
	"flag"
	"fmt"
	"strings"
)

var n = flag.Bool("n", false, "omit trailing new line")
var sep = flag.String("sep", " ", "separator")

func main() {
	flag.Parse()
	fmt.Print(strings.Join(flag.Args(), *sep))
	if *n == false && len(flag.Args()) > 0 {
		fmt.Println()
	}
}
