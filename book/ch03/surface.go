package ch03

import "math"

const (
	Width, Height = 600, 300            // canvas size in pixels
	Cells         = 100                 // number of grid cells
	xyrange       = 30.0                // axis ranges (-xrange..+xrange)
	xyscale       = Width / 2 / xyrange // pixels per x or y unit
	zscale        = Height * 0.4        // pixels per z unit
	angle         = math.Pi / 6         // angle of x, y axis (-30)
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle)

func Corner(i, j int) (float64, float64) {
	// Find pint (x,y) at corner of cell (i,j)
	x := xyrange * (float64(i)/Cells - 0.5)
	y := xyrange * (float64(j)/Cells - 0.5)

	// Compute surface height z.
	z := f(x, y)

	// Project (x,y,z) isometrically onto 2-D SVG canvas (sx,sy)
	sx := Width/2 + (x-y)*cos30*xyscale
	sy := Height/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy
}

func f(x, y float64) float64 {
	r := math.Hypot(x, y) // distance from (0,0)
	return math.Sin(r) / r
}
