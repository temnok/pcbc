package xunpu

import (
	"temnok/pcbc/eda/lib"
	"temnok/pcbc/path"
)

var SwitchTS1088 = &lib.Component{
	Pads: path.Rect(1.35, 1.8).
		Clone(2, 4.15, 0),
	Marks: path.Strokes{
		0.1: path.Paths{
			path.Rect(4, 2.9),
		},
	},
}
