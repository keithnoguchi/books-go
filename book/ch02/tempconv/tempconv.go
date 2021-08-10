// tempconv converts a temperature between Fahrenheit and Celsius
package tempconv

import "fmt"

type Fahrenheit float64
type Celsius float64

const (
	AbsoluteZeroC Celsius = -273.15
	FreezingC     Celsius = 0
	BoilingC      Celsius = 100
)

func (f Fahrenheit) String() string {
	return fmt.Sprintf("%gF", f)
}

func (f Fahrenheit) ToCelsius() Celsius {
	return Celsius((f - 32) * 5 / 9)
}

func (c Celsius) String() string {
	return fmt.Sprintf("%gC", c)
}

func (c Celsius) ToFahrenheit() Fahrenheit {
	return Fahrenheit(c*9/5 + 32)
}
