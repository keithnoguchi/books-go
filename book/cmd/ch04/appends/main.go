// appendint appends integer to the slice of integer.
package main

import (
	"fmt"

	appends "book/ch04/appends/v1"
)

func main() {
	var x, y []int
	for i := 0; i < 10; i++ {
		y = appends.AppendInt(x, i)
		fmt.Printf("%d  len=%-2d cap=%-2d\t%v\n", i, len(y), cap(y), y)
		x = y
	}
}
