package holder

import (
	"temnok/pcbc/eda/lib"
	"temnok/pcbc/path"
)

var LIR1254 = &lib.Component{
	Pads: path.Rect(2.5, 4.5).Clone(2, 15, 0).Append(
		path.Paths{path.Circle(5)},
	),

	Marks: path.Strokes{
		0.1: path.Paths{
			path.Circle(12),
		},
	},
}
