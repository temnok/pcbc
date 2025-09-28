// Copyright © 2025 Alex Temnok. All rights reserved.

package worldsemi

import (
	"temnok/pcbc/eda"
	"temnok/pcbc/path"
	"temnok/pcbc/transform"
)

var (
	padPair = path.Rect(0.7, 0.7).Clone(2, 0, (0.7 + 0.4)).Transform(transform.Move((0.7+1.13)/2, 0))

	WS2812B_2020 = &eda.Component{
		Pads: append(append(
			path.Paths{nil},
			padPair...,
		), padPair.Transform(transform.RotateDegrees(-180))...),

		Marks: path.Paths{
			path.Rect(2.4, 2.2),
			path.Rect(0.3, 0.3).Transform(transform.Move(0, 0.4)),
			path.Rect(0.05, 0.05).Transform(transform.Move(-0.3, -0.4)),
			path.Rect(0.05, 0.05).Transform(transform.Move(0, -0.4)),
			path.Rect(0.05, 0.05).Transform(transform.Move(0.3, -0.4)),
		},
	}
)
