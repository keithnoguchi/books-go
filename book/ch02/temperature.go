package ch02

type Celsius float64
type Fahrenheit float64

const (
	AbsoluteZeroC Celsius = -273.15
	FreezingC     Celsius = 0
	BoilingC      Celsius = 100
	FreezingF     Fahrenheit = 32
	BoilingF      Fahrenheit = 212
)

func FtoC(f Fahrenheit) Celsius {
	return Celsius((f - 32) * 5 / 9)
}

func CtoF(c Celsius) Fahrenheit {
	return Fahrenheit(c * 9 / 5 + 32)
}
