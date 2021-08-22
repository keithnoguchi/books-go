package main

import (
	"fmt"

	"book/ch06"
)

func main() {
	p := ch06.Point{1, 2}
	q := ch06.Point{2, 4}
	fmt.Println(p.Distance(q))
	path := ch06.Path{
		{1, 1},
		{5, 1},
		{5, 4},
		{1, 1},
	}
	fmt.Println(path.Distance())
}
