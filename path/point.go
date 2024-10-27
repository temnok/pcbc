package path

import (
	"math"
	"temnok/pcbc/transform"
)

type Point struct {
	X, Y float64
}

func (p Point) Round() Point {
	return Point{math.Round(p.X), math.Round(p.Y)}
}

func (p Point) RoundXY() (x, y int) {
	return int(math.Round(p.X)), int(math.Round(p.Y))
}

func (p Point) Apply(t transform.Transform) Point {
	x, y := t.Apply(p.X, p.Y)
	return Point{X: x, Y: y}
}
