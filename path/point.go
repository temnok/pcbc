// Copyright © 2025 Alex Temnok. All rights reserved.

package path

import (
	"github.com/temnok/pcbc/transform"
	"math"
)

type Point struct {
	X, Y float64
}

func (p Point) XY() (x, y float64) {
	return p.X, p.Y
}

func (p Point) Round() Point {
	return Point{math.Round(p.X), math.Round(p.Y)}
}

func (p Point) RoundXY() (x, y int) {
	return int(math.Round(p.X)), int(math.Round(p.Y))
}

func (p Point) Apply(t transform.T) Point {
	x, y := t.Apply(p.X, p.Y)
	return Point{X: x, Y: y}
}

func (p Point) Move(dx, dy float64) Point {
	return Point{X: p.X + dx, Y: p.Y + dy}
}
