package pkg

import (
	"temnok/lab/contour"
	"temnok/lab/geom"
)

func qfn17pads() [][]geom.XY {
	const (
		b  = 0.24
		D2 = 1.7
		E  = 3.0
		E2 = 1.7
		e  = 0.5
		L  = 0.3
		Lp = L * 2
	)

	var pads [][]geom.XY

	row := padRow(4, e, b, Lp)
	pads = geom.MoveXY(-E/2, 0).RotateD(90).PointsAll(row)
	pads = append(pads, geom.MoveXY(0, E/2).PointsAll(row)...)
	pads = append(pads, geom.MoveXY(E/2, 0).RotateD(90).PointsAll(row)...)
	pads = append(pads, geom.MoveXY(0, -E/2).PointsAll(row)...)

	pads = append(pads, contour.Rect(D2, E2))

	return pads
}
