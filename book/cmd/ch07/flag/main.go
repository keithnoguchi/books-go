// SPDX-License-Identifier: GPL-2.0
package main

import (
	"flag"
	"fmt"

	"book/ch07"
)

var temp = ch07.CelciousFlag("temp", 20.0, "the temperature")

func main() {
	flag.Parse()
	fmt.Printf("%s temperature\n", *temp)
}
