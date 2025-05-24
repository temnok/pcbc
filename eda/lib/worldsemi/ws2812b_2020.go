// Copyright © 2025 Alex Temnok. All rights reserved.

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
			pad2020.Transform(transform.Move(pad2020_X, -pad2020_Y)),
			pad2020.Transform(transform.Move(pad2020_X, pad2020_Y)),
			pad2020.Transform(transform.Move(-pad2020_X, pad2020_Y)),
			pad2020.Transform(transform.Move(-pad2020_X, -pad2020_Y)),
		},

		Marks: path.Paths{
			path.Rect(2.20, 2),
			path.Rect(0.3, 0.3).Transform(transform.Move(0, 0.4)),
			path.Rect(0.05, 0.05).Transform(transform.Move(-0.3, -0.4)),
			path.Rect(0.05, 0.05).Transform(transform.Move(0, -0.4)),
			path.Rect(0.05, 0.05).Transform(transform.Move(0.3, -0.4)),
		},
	}
)
