package contour

import "temnok/lab/geom"

func Size(c []geom.XY) geom.XY {
	lo, hi := c[0], c[0]

	for _, p := range c[1:] {
		lo.X, lo.Y = min(lo.X, p.X), min(lo.Y, p.Y)
		hi.X, hi.Y = max(hi.X, p.X), max(hi.Y, p.Y)
	}

	return geom.XY{hi.X - lo.X, hi.Y - lo.Y}
}
