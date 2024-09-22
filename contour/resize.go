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
	kx := (ma.X - mi.X + delta) / (ma.X - mi.X)
	ky := (ma.Y - mi.Y + delta) / (ma.Y - mi.Y)
	c := geom.XY{(mi.X + ma.X) / 2, (mi.Y + ma.Y) / 2}

	res := make([]geom.XY, len(contour))

	for i, p := range contour {
		res[i] = geom.XY{(1-kx)*c.X + kx*p.X, (1-ky)*c.Y + ky*p.Y}
	}

	return res
}
