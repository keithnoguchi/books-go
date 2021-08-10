package main

import "fmt"

func main() {
	var a []int
	fmt.Printf("%p\n", a)
	a = nil
	fmt.Printf("%p\n", a)
	a = []int(nil)
	fmt.Printf("%p\n", a)
	a = []int{}
	fmt.Printf("%p\n", a)
}
