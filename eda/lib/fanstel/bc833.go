package fanstel

import (
	"temnok/pcbc/eda/lib"
	"temnok/pcbc/geom"
	"temnok/pcbc/path"
)

var (
	pad = path.Circle(0.5)

	BC833 = &lib.Component{
		Pads: path.Paths{
			0:  nil,
			1:  pad.Transform(geom.MoveXY(-4.5, 4.8)),
			2:  pad.Transform(geom.MoveXY(-4.5, 3.8)),
			3:  pad.Transform(geom.MoveXY(-4.5, 2.8)),
			4:  pad.Transform(geom.MoveXY(-4.5, 1.8)),
			5:  pad.Transform(geom.MoveXY(-4.5, 0.8)),
			6:  pad.Transform(geom.MoveXY(-3.5, 3.8)),
			7:  pad.Transform(geom.MoveXY(-3.5, 2.8)),
			8:  pad.Transform(geom.MoveXY(-3.5, 1.8)),
			9:  pad.Transform(geom.MoveXY(-0.5, -0.2)),
			10: pad.Transform(geom.MoveXY(0.5, -0.2)),
			11: pad.Transform(geom.MoveXY(1.5, -0.2)),
			12: pad.Transform(geom.MoveXY(2.5, -0.2)),
			13: pad.Transform(geom.MoveXY(3.5, -0.2)),
			14: pad.Transform(geom.MoveXY(4.5, -0.2)),
			15: pad.Transform(geom.MoveXY(4.5, 1.8)),
			16: pad.Transform(geom.MoveXY(4.5, 2.8)),
			17: pad.Transform(geom.MoveXY(4.5, 3.8)),
			18: pad.Transform(geom.MoveXY(4.5, 4.8)),
			19: pad.Transform(geom.MoveXY(3.5, 1.8)),
			20: pad.Transform(geom.MoveXY(3.5, 2.8)),
			21: pad.Transform(geom.MoveXY(3.5, 3.8)),
			22: pad.Transform(geom.MoveXY(3.5, 4.8)),
			23: pad.Transform(geom.MoveXY(2.5, 4.8)),
			24: pad.Transform(geom.MoveXY(1.5, 4.8)),
			25: pad.Transform(geom.MoveXY(0.5, 2.8)),
			26: pad.Transform(geom.MoveXY(0.5, 1.8)),
			27: pad.Transform(geom.MoveXY(-0.5, 2.8)),
			28: pad.Transform(geom.MoveXY(-0.5, 1.8)),
		}.Transform(geom.MoveXY(0, 2-14.3/2)),

		Marks: path.Strokes{
			0.2: {
				path.Rect(10, 14.3),
			},
		},
	}
)
