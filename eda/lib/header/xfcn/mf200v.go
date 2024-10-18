package xfcn

import (
	"temnok/pcbc/eda/lib"
	"temnok/pcbc/path"
)

func MF200V(n int) *lib.Component {
	const w, h = 2, 2.2

	return &lib.Component{
		Cuts: path.Circle(0.6).Clone(n, w, 0),
		Pads: path.Circle(1.2).Clone(n, w, 0),
		Marks: path.Strokes{
			0.1: path.CutRect(w, h, 0.3).Clone(n, w, 0),
		},
	}
}
