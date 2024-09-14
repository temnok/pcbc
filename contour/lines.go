package contour

import "temnok/lab/geom"

func Lines(points []geom.XY) []geom.XY {
	if len(points) == 0 {
		return nil
	}

	res := []geom.XY{points[0]}

	for i, p := range points[1:] {
		res = append(res, res[i], p, p)
	}

	return res
}
