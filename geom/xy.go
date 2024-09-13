package geom

import "math"

type XY struct {
	X, Y float64
}

func (xy XY) Round() XY {
	return XY{math.Round(xy.X), math.Round(xy.Y)}
}
