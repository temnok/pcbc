package mph100imp40f

import (
	"temnok/pcbc/eda"
	"temnok/pcbc/path"
)

func G_V_SP(n int) *eda.Component {
	const step = 2.54

	return &eda.Component{
		Pads:  path.Circle(1.8).Clone(n, step, 0),
		Marks: path.CutRect(step, step, 0.3).Clone(n, step, 0),
	}
}
