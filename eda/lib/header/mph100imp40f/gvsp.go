package mph100imp40f

import (
	"temnok/lab/eda/lib"
	"temnok/lab/path"
)

var G_V_SP_9 = G_V_SP(9)

func G_V_SP(n int) *lib.Component {
	const tenth = 2.54

	pads := path.Circle(1.5).Clone(n, tenth, 0)
	return &lib.Component{
		Pads:     pads,
		Openings: pads,
	}
}
