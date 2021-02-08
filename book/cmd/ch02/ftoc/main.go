// ftoc prints two Fahrenheit-to-Celcius conversions.
package main

import (
	"fmt"

	"book/ch02"
)

func main() {
	const freezingF, boilingF = 32.0, 212.0
	fmt.Printf("%gF = %gC\n", freezingF, ch02.FtoC(freezingF))
	fmt.Printf("%gF = %gC\n", boilingF, ch02.FtoC(boilingF))
	fmt.Printf("%gF = %gC\n", ch02.CtoF(ch02.FtoC(freezingF)), ch02.FtoC(freezingF))
	fmt.Printf("%gF = %gC\n", ch02.CtoF(ch02.FtoC(boilingF)), ch02.FtoC(boilingF))
}
