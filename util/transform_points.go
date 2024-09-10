package util

import (
	"temnok/lab/bezier"
	"temnok/lab/t2d"
)

func TransformPoints(t t2d.Transform, points []bezier.Point) []bezier.Point {
	res := make([]bezier.Point, len(points))

	for i, p := range points {
		r := t.Point(t2d.Vector{p.X, p.Y})
		res[i].X, res[i].Y = r[0], r[1]
	}

	return res
}

func TransformAllPoints(t t2d.Transform, allPoints [][]bezier.Point) [][]bezier.Point {
	res := make([][]bezier.Point, len(allPoints))

	for i, points := range allPoints {
		res[i] = TransformPoints(t, points)
	}

	return res
}
