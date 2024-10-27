package mph100imp40f

import (
	"temnok/pcbc/eda"
	"temnok/pcbc/path"
)

var (
	G_V_SP_x2  = G_V_SP(2)
	G_V_SP_x3  = G_V_SP(3)
	G_V_SP_x4  = G_V_SP(4)
	G_V_SP_x6  = G_V_SP(6)
	G_V_SP_x8  = G_V_SP(8)
	G_V_SP_x9  = G_V_SP(9)
	G_V_SP_x10 = G_V_SP(10)
	G_V_SP_x11 = G_V_SP(11)
)

func G_V_SP(n int) *eda.Component {
	const step = 2.54

	return &eda.Component{
		Pads: path.Circle(1.8).Clone(n, step, 0),
		Marks: path.Strokes{
			0.1: path.CutRect(step, step, 0.3).Clone(n, step, 0),
		},
	}
}
