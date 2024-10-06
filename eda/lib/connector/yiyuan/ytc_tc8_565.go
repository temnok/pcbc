package yiyuan

import (
	"temnok/lab/eda/lib"
	"temnok/lab/geom"
	"temnok/lab/path"
)

var (
	pads = path.Rect(0.5, 1.6).Clone(8, 0.8, 0).Transform(geom.MoveXY(0, -0.88-0.3))

	holes = path.RoundRect(0.6, 1.6, 0.3).Clone(2, 8.64, 0).Transform(geom.MoveXY(0, -0.3))

	YTC_TC8_565 = &lib.Component{
		Pads:     pads,
		Openings: pads,
		//Holes:    holes,
		//Openings: append(pads, holes...),
		Marks: path.Strokes{
			0.1: path.Paths{
				path.RoundRect(9, 3.16, 1),
			},
		},
	}
)
