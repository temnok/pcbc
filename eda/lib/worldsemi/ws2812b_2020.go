package worldsemi

import (
	"temnok/pcbc/eda"
	"temnok/pcbc/path"
	"temnok/pcbc/transform"
)

var (
	pad2020              = path.Rect(0.8, 0.8)
	pad2020_X, pad2020_Y = 0.915, 0.55

	WS2812B_2020 = &eda.Component{
		Pads: path.Paths{
			1: pad2020.Apply(transform.Move(pad2020_X, -pad2020_Y)),
			2: pad2020.Apply(transform.Move(pad2020_X, pad2020_Y)),
			3: pad2020.Apply(transform.Move(-pad2020_X, pad2020_Y)),
			4: pad2020.Apply(transform.Move(-pad2020_X, -pad2020_Y)),
		},

		Marks: path.Paths{
			path.Rect(2.20, 2),
			path.Rect(0.3, 0.3).Apply(transform.Move(0, 0.4)),
			path.Rect(0.05, 0.05).Apply(transform.Move(-0.3, -0.4)),
			path.Rect(0.05, 0.05).Apply(transform.Move(0, -0.4)),
			path.Rect(0.05, 0.05).Apply(transform.Move(0.3, -0.4)),
		},
	}
)
