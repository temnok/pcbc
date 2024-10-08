package sop

import (
	"temnok/lab/eda/lib"
	"temnok/lab/geom"
	"temnok/lab/path"
)

var (
	pad = path.RoundRect(0.5, 1.0, 0.25)

	bottomPads = pad.Clone(4, 1.27, 0).Transform(geom.MoveXY(0, -2.75))

	topPads = bottomPads.Transform(geom.RotateD(180))

	pads = append(bottomPads, topPads...)

	SOP8 = &lib.Component{
		Pads: pads,
		Marks: path.Strokes{
			0.1: path.Paths{
				path.Rect(5, 4),
				path.Circle(0.3).Transform(geom.MoveXY(-2.4, -2.3)),
			},
		},
	}
)
