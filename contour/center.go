package contour

import "temnok/lab/geom"

func Center(contour []geom.XY) geom.XY {
	x, y, n := 0.0, 0.0, 0.0

	for i := 0; i < len(contour); i += 3 {
		x += contour[i].X
		y += contour[i].Y
		n++
	}

	return geom.XY{x / n, y / n}
}
