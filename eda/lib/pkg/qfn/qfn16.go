package qfn

import (
	"temnok/lab/eda/lib"
	"temnok/lab/geom"
	"temnok/lab/path"
)

var QFN16G *lib.Component

func init() {
	pad := path.RoundRect(0.7, 0.25, 0.12)

	col := pad.Clone(4, 0, -0.5).Transform(geom.MoveXY(-1.5, 0))

	pads := path.Paths{}
	for a := 0.0; a < 360; a += 90 {
		pads = append(pads, col.Transform(geom.RotateD(a))...)
	}

	pads = append(pads, path.Rect(1.2, 1.2))

	QFN16G = &lib.Component{
		Pads: pads,
		Marks: path.Strokes{
			0.1: path.Paths{
				path.Rect(3.1, 3.1),
				path.Circle(0.3).Transform(geom.MoveXY(-1.8, 1.4)),
			},
		},
	}
}
