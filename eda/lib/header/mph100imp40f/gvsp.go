package mph100imp40f

import (
	"temnok/pcbc/eda/lib"
	"temnok/pcbc/path"
)

var (
	G_V_SP_x2 = G_V_SP(2)
	G_V_SP_x4 = G_V_SP(4)
	G_V_SP_x6 = G_V_SP(6)
	G_V_SP_x8 = G_V_SP(8)
	G_V_SP_x9 = G_V_SP(9)
)

func G_V_SP(n int) *lib.Component {
	const tenth = 2.54

	return &lib.Component{
		Pads: path.Circle(1.5).Clone(n, tenth, 0),
		Marks: path.Strokes{
			0.1: path.CutRect(tenth, tenth, 0.3).Clone(n, tenth, 0),
		},
	}
}
