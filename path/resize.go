package path

import "temnok/pcbc/geom"

func (path Path) Resize(delta float64) Path {
	mi, ma := path.minMax()
	kx := (ma.X - mi.X + delta) / (ma.X - mi.X)
	ky := (ma.Y - mi.Y + delta) / (ma.Y - mi.Y)
	center := geom.XY{(mi.X + ma.X) / 2, (mi.Y + ma.Y) / 2}

	res := make([]geom.XY, len(path))

	for i, p := range path {
		res[i] = geom.XY{(1-kx)*center.X + kx*p.X, (1-ky)*center.Y + ky*p.Y}
	}

	return res
}

func (path Path) minMax() (mi, ma geom.XY) {
	if len(path) > 0 {
		mi, ma = path[0], path[0]

		for _, p := range path {
			mi.X = min(mi.X, p.X)
			mi.Y = min(mi.Y, p.Y)
			ma.X = max(ma.X, p.X)
			ma.Y = max(ma.Y, p.Y)
		}
	}

	return
}
