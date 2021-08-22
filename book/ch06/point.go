package ch06

import "math"

type Point struct{ X, Y float64 }

func (p *Point) Distance(q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

type Path []Point

func (p *Path) Distance() float64 {
	var sum float64
	for i := range *p {
		if i > 0 {
			sum += (*p)[i].Distance((*p)[i-1])
		}
	}
	return sum
}
