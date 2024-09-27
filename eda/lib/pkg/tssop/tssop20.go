package tssop

import (
	"temnok/lab/geom"
	"temnok/lab/path"
)

func TSSOP20() path.Paths {
	row := path.Rect(0.3, 1.35).
		Clone(10, 0.65, 0).
		Transform(geom.MoveXY(0, -2.925))

	return append(
		row,
		row.Transform(geom.RotateD(180))...,
	)
}
