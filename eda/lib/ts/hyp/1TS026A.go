package hyp

import (
	"temnok/pcbc/eda"
	"temnok/pcbc/path"
	"temnok/pcbc/transform"
)

var Switch1TS026A = &eda.Component{
	Pads: path.Rect(0.5, 0.55).
		CloneRowsCols(2, 2, path.Point{X: 1.3 - 0.5, Y: 3.2 - 0.55}).
		Apply(transform.Scale(1, -1).Rotate(90)),

	Marks: path.Paths{
		path.Rect(2.6, 1.6),
	},
}
