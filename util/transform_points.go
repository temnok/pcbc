package util

import (
	"temnok/lab/twod"
)

func TransformPoints(t twod.Transform, points []twod.Coord) []twod.Coord {
	res := make([]twod.Coord, len(points))

	for i, p := range points {
		r := t.Point(twod.Coord{X: p.X, Y: p.Y})
		res[i].X, res[i].Y = r.X, r.Y
	}

	return res
}

func TransformAllPoints(t twod.Transform, allPoints [][]twod.Coord) [][]twod.Coord {
	res := make([][]twod.Coord, len(allPoints))

	for i, points := range allPoints {
		res[i] = TransformPoints(t, points)
	}

	return res
}
