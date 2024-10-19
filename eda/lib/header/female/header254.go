package female

import (
	"temnok/pcbc/eda/lib"
	"temnok/pcbc/path"
)

func Header254(n int) *lib.Component {
	const w, h = 2.54, 2.54

	return &lib.Component{
		Cuts: path.Circle(0.7).Clone(n, w, 0),
		Pads: path.Circle(1.8).Clone(n, w, 0),
		Marks: path.Strokes{
			0.1: path.CutRect(w, h, 0.3).Clone(n, w, 0),
		},
	}
}
