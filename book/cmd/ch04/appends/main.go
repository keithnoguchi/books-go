// appendint appends integer to the slice of integer.
package main

import (
	"fmt"

	appends "book/ch04/appends/v2"
)

func main() {
	var x, y []int
	for i := 0; i < 10; i += 2 {
		y = appends.AppendInt(x, []int{i, i + 1}...)
		fmt.Printf("%d  len=%-2d cap=%-2d\t%v\n", i, len(y), cap(y), y)
		x = y
	}
}
