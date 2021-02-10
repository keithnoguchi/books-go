// comma adds comma, ',', to the financial numbers.
package main

import (
	"fmt"
	"os"

	"book/ch03"
)

func main() {
	var input string
	if len(os.Args) <= 1 {
		input = "300"
	} else {
		input = os.Args[1]
	}
	fmt.Println(ch03.Comma(input))
}
