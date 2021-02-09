package mandelbrot

import (
	"image/color"
	"math/cmplx"
)

const (
	Xmin, Ymin, Xmax, Ymax = -2, -2, +2, +2
	Width, Height          = 1024, 1024
)

func Mandelbrot(z complex128) color.Color {
	const iterations = 200
	const contrast = 15

	var v complex128
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			return color.Gray{255 - contrast*n}
		}
	}
	return color.Black
}
