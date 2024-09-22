package contour

import "temnok/lab/geom"

func minMax(contour []geom.XY) (mi, ma geom.XY) {
	mi, ma = contour[0], contour[0]

	for _, p := range contour {
		mi.X = min(mi.X, p.X)
		mi.Y = min(mi.Y, p.Y)
		ma.X = max(ma.X, p.X)
		ma.Y = max(ma.Y, p.Y)
	}

	return
}

func Resize(contour []geom.XY, delta float64) []geom.XY {
	mi, ma := minMax(contour)
	size := max(ma.X-mi.X, ma.Y-mi.Y)
	k := (size + delta) / size
	c := geom.XY{(mi.X + ma.X) / 2, (mi.Y + ma.Y) / 2}

	res := make([]geom.XY, len(contour))

	for i, p := range contour {
		res[i] = mix(c, p, k)
	}

	return res
}

func mix(a, b geom.XY, k float64) geom.XY {
	return geom.XY{(1-k)*a.X + k*b.X, (1-k)*a.Y + k*b.Y}
}
