// bit: bit operations
package main

import "fmt"

func main() {
	var x uint8 = 1<<1 | 1<<5
	var y uint8 = 1<<1 | 1<<2

	fmt.Printf("x\t%08b\n", x)
	fmt.Printf("y\t%08b\n", y)
	fmt.Printf("x&y\t%08b\n", x&y)
	fmt.Printf("x|y\t%08b\n", x|y)
	fmt.Printf("x^y\t%08b\n", x^y)
	fmt.Printf("x&^y\t%08b\n", x&^y)

	// print out the x's set bit position
	for i := uint(0); i < 8; i++ {
		if x&(1<<i) != 0 {
			fmt.Println(i)
		}
	}
	fmt.Printf("x<<1\t%08b\n", x<<1)
	fmt.Printf("x>>1\t%08b\n", x>>1)
}
