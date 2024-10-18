package hyp

import (
	"temnok/pcbc/eda/lib"
	"temnok/pcbc/geom"
	"temnok/pcbc/path"
	"temnok/pcbc/transform"
)

var Switch1TS026A = &lib.Component{
	Pads: path.Rect(0.5, 0.55).
		CloneRowsCols(2, 2, geom.XY{1.3 - 0.5, 3.2 - 0.55}).
		Apply(transform.Scale(1, -1).Rotate(90)),
	Marks: path.Strokes{
		0.1: path.Paths{
			path.Rect(2.6, 1.6),
		},
	},
}
