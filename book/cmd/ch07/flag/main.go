// SPDX-License-Identifier: GPL-2.0
package main

import (
	"flag"
	"fmt"
	"time"
)

var period = flag.Duration("period", 1*time.Second, "sleep period")

func main() {
	flag.Parse()
	fmt.Printf("sleep for %s\n", *period)
	time.Sleep(*period)
	fmt.Println()
}
