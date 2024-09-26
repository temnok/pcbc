package contour

import (
	"temnok/lab/geom"
	"temnok/lab/path"
)

func Lines(points []geom.XY) path.Path {
	if len(points) == 0 {
		return nil
	}

	res := path.Path{points[0]}

	for _, p := range points[1:] {
		res = append(res, res[len(res)-1], p, p)
	}

	return res
}
