package geom

import "math"

type XY struct {
	X, Y float64
}

func (xy XY) Round() XY {
	return XY{math.Round(xy.X), math.Round(xy.Y)}
}

func (xy XY) Ints() (x, y int) {
	return int(math.Round(xy.X)), int(math.Round(xy.Y))
}
