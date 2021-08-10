// rotate rotates the slice for nth to right
package ch04

func Rotate(a []interface{}, nth int) {
	Reverse(a[:nth])
	Reverse(a[nth:])
	Reverse(a)
}
