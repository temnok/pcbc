package fanstel

import (
	"temnok/pcbc/eda"
	"temnok/pcbc/path"
	"temnok/pcbc/transform"
)

var (
	pad = path.Circle(0.6)

	BC833 = &eda.Component{
		Pads: path.Paths{
			0:  nil,
			1:  pad.Apply(transform.Move(-4.5, 4.8)),
			2:  pad.Apply(transform.Move(-4.5, 3.8)),
			3:  pad.Apply(transform.Move(-4.5, 2.8)),
			4:  pad.Apply(transform.Move(-4.5, 1.8)),
			5:  pad.Apply(transform.Move(-4.5, 0.8)),
			6:  pad.Apply(transform.Move(-3.5, 3.8)),
			7:  pad.Apply(transform.Move(-3.5, 2.8)),
			8:  pad.Apply(transform.Move(-3.5, 1.8)),
			9:  pad.Apply(transform.Move(-0.5, -0.2)),
			10: pad.Apply(transform.Move(0.5, -0.2)),
			11: pad.Apply(transform.Move(1.5, -0.2)),
			12: pad.Apply(transform.Move(2.5, -0.2)),
			13: pad.Apply(transform.Move(3.5, -0.2)),
			14: pad.Apply(transform.Move(4.5, -0.2)),
			15: pad.Apply(transform.Move(4.5, 1.8)),
			16: pad.Apply(transform.Move(4.5, 2.8)),
			17: pad.Apply(transform.Move(4.5, 3.8)),
			18: pad.Apply(transform.Move(4.5, 4.8)),
			19: pad.Apply(transform.Move(3.5, 1.8)),
			20: pad.Apply(transform.Move(3.5, 2.8)),
			21: pad.Apply(transform.Move(3.5, 3.8)),
			22: pad.Apply(transform.Move(3.5, 4.8)),
			23: pad.Apply(transform.Move(2.5, 4.8)),
			24: pad.Apply(transform.Move(1.5, 4.8)),
			25: pad.Apply(transform.Move(0.5, 2.8)),
			26: pad.Apply(transform.Move(0.5, 1.8)),
			27: pad.Apply(transform.Move(-0.5, 2.8)),
			28: pad.Apply(transform.Move(-0.5, 1.8)),
		}.Apply(transform.Move(0.15, 1.75-14.3/2)),

		MarkStrokes: path.Strokes{
			0.1: {
				path.Rect(10, 14.3),
			},
		},
	}
)
