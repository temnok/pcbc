package xunpu

import (
	"temnok/pcbc/eda"
	"temnok/pcbc/path"
)

var SwitchTS1088 = &eda.Component{
	Pads: path.Paths{path.Rect(1.35, 1.8)}.
		Clone(2, 4.15, 0),

	Marks: path.Paths{
		path.Rect(4, 2.9),
	},
}
