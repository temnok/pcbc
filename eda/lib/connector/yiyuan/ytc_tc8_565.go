package yiyuan

import (
	"temnok/pcbc/eda/lib"
	"temnok/pcbc/path"
	"temnok/pcbc/transform"
)

var (
	YTC_TC8_565 = &lib.Component{
		Pads: path.Rect(0.5, 1.6).Clone(8, 0.8, 0).Apply(transform.Move(0, -0.88-0.3)),
		Marks: path.Strokes{
			0.1: path.Paths{
				path.RoundRect(9, 3.16, 1),
			},
		},
	}
)
