package hyp

import (
	"temnok/pcbc/eda/lib"
	"temnok/pcbc/geom"
	"temnok/pcbc/path"
)

var Switch1TS026A = &lib.Component{
	Pads: path.Rect(0.5, 0.55).
		CloneColsRows(2, 2, geom.XY{1.3 - 0.5, 3.2 - 0.55}).
		Transform(geom.RotateD(90)),
	Marks: path.Strokes{
		0.1: path.Paths{
			path.Rect(2.6, 1.6),
		},
	},
}
