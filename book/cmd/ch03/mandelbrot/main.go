// mandelbrot emits a PNG image of the Mandelbrot fractal.
package main

import (
	"image"
	"image/png"
	"os"

	ch03 "book/ch03/mandelbrot"
)

func main() {
	img := image.NewRGBA(image.Rect(0, 0, ch03.Width, ch03.Height))
	for py := 0; py < ch03.Height; py++ {
		y := float64(py)/ch03.Height*(ch03.Ymax-ch03.Ymin) + ch03.Ymin
		for px := 0; px < ch03.Width; px++ {
			x := float64(px)/ch03.Width*(ch03.Xmax-ch03.Xmin) +
				ch03.Xmin
			z := complex(x, y)
			// Image point (px, py) represents compex value z.
			img.Set(px, py, ch03.Mandelbrot(z))
		}
	}
	png.Encode(os.Stdout, img)
}
