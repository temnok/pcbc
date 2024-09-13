package util

import "temnok/lab/geom"

func TransformPoints(t geom.Transform, points []geom.XY) []geom.XY {
	res := make([]geom.XY, len(points))

	for i, p := range points {
		r := t.Point(geom.XY{X: p.X, Y: p.Y})
		res[i].X, res[i].Y = r.X, r.Y
	}

	return res
}

func TransformAllPoints(t geom.Transform, allPoints [][]geom.XY) [][]geom.XY {
	res := make([][]geom.XY, len(allPoints))

	for i, points := range allPoints {
		res[i] = TransformPoints(t, points)
	}

	return res
}
