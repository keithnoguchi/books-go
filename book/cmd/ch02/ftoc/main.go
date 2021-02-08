// ftoc prints two Fahrenheit-to-Celcius conversions.
package main

import (
	"fmt"

	"book/ch02"
)

func main() {
	const freezingF, boilingF = ch02.FreezingF, ch02.BoilingF
	fmt.Printf("%v = %v\n", freezingF, freezingF.Celsius())
	fmt.Printf("%v = %v\n", boilingF, boilingF.Celsius())
	fmt.Printf("%v = %v\n", freezingF.Celsius().Fahrenheit(), freezingF.Celsius())
	fmt.Printf("%v = %v\n", boilingF.Celsius().Fahrenheit(), boilingF.Celsius())
}
