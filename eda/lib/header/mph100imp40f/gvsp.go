package mph100imp40f

import (
	"temnok/lab/eda/lib"
	"temnok/lab/path"
)

var (
	G_V_SP_x4 = G_V_SP(4)
	G_V_SP_x9 = G_V_SP(9)
)

func G_V_SP(n int) *lib.Component {
	const tenth = 2.54

	pads := path.Circle(1.5).Clone(n, tenth, 0)

	marks := path.CutRect(tenth, tenth, 0.3).Clone(n, tenth, 0)

	return &lib.Component{
		Pads:     pads,
		Openings: pads,
		Marks: path.Strokes{
			0.1: marks,
		},
	}
}
