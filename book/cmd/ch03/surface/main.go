// surface computes an SVG rendering of a 3-D surface function.
package main

import (
	"fmt"

	"book/ch03"
)

func main() {
	fmt.Printf("<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 07' "+
		"width='%d' height='%d'>", ch03.Width, ch03.Height)
	for i := 0; i < ch03.Cells; i++ {
		for j := 0; j < ch03.Cells; j++ {
			ax, ay := ch03.Corner(i+1, j)
			bx, by := ch03.Corner(i, j)
			cx, cy := ch03.Corner(i, j+1)
			dx, dy := ch03.Corner(i+1, j+1)
			fmt.Printf("<polygon points='%g,%g %g,%g %g,%g %g,%g'/>\n",
				ax, ay, bx, by, cx, cy, dx, dy)
		}
	}
	fmt.Println("</svg>")
}
