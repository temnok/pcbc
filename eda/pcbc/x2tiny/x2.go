package x2tiny

import (
	"temnok/pcbc/eda/lib"
	"temnok/pcbc/eda/lib/header/greenconn"
	"temnok/pcbc/eda/lib/pkg/smd"
	"temnok/pcbc/eda/pcbc"
	"temnok/pcbc/path"
	"temnok/pcbc/transform"
)

var X2Base = &lib.Component{
	Cuts: path.Paths{
		path.RoundRect(4.75, 6.75, 1),
	},

	Components: lib.Components{
		smd.I0603.Arrange(transform.Rotate(90).Move(-1.5, -0.5)),
		greenconn.CSCC118(2, false, []string{"", "K20"}).Arrange(transform.Move(0, -1.8)),
		pcbc.MountHole.Arrange(transform.Move(0.5, 1.5)),
	},
}
