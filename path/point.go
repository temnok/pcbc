package path

import "math"

type Point struct {
	X, Y float64
}

func (p Point) Round() Point {
	return Point{math.Round(p.X), math.Round(p.Y)}
}

func (p Point) RoundXY() (x, y int) {
	return int(math.Round(p.X)), int(math.Round(p.Y))
}
