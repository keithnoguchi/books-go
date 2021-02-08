// cf converts it's numeric argument to Celsius and Fahrenheit.
package main

import (
	"fmt"
	"os"
	"strconv"

	"book/ch02"
)

func main() {
	for _, arg := range os.Args[1:] {
		t, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "cf: %v\n", err)
			os.Exit(1)
		}
		f := ch02.Fahrenheit(t)
		c := ch02.Celsius(t)
		fmt.Printf("%s = %s, %s = %s\n",
			f, f.Celsius(), c, c.Fahrenheit())
	}
}
