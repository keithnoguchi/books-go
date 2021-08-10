package main

import (
	"fmt"

	"book/ch04"
)

func main() {
	var s ch04.Stack

	s.Push(interface{}(1))
	s.Push(interface{}(2))
	s.Push(interface{}(3))
	s.Push(interface{}('a'))
	s.Push(interface{}("a"))
	fmt.Printf("len(s)=%[1]T,len(s)=%[1]d, %v\n", len(s), s)

	for len(s) > 0 {
		fmt.Println(s.Pop())
	}
}
