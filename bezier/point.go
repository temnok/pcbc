package bezier

import "math"

type Point struct{ X, Y float64 }

//go:inline
func (a Point) Mix(b Point, t float64) Point {
	return Point{a.X*(1-t) + b.X*t, a.Y*(1-t) + b.Y*t}
}

func (a Point) Round() Point {
	return Point{math.Round(a.X), math.Round(a.Y)}
}
