package twod

import "math"

type Coord struct {
	X, Y float64
}

func (c Coord) Round() Coord {
	return Coord{
		X: math.Round(c.X),
		Y: math.Round(c.Y),
	}
}
