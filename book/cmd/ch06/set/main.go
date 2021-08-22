package main

import (
	"fmt"

	"book/ch06"
)

func main() {
	var s ch06.IntSet
	s.Add(99)
	fmt.Println(s.Has(99))
	fmt.Println(s.Has(100))
	fmt.Println(s.Has(98))
	s.Add(64)
	fmt.Println(s.Has(64))
	fmt.Println(s.Has(63))
	fmt.Println(s.Has(65))
	fmt.Println(&s)
}
