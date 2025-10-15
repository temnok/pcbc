// Copyright © 2025 Alex Temnok. All rights reserved.

package worldsemi

import (
	"github.com/temnok/pcbc/eda"
	"github.com/temnok/pcbc/path"
	"github.com/temnok/pcbc/transform"
)

var (
	padPair = path.RoundRect(0.7, 0.7, 0.2).CloneXY(2, 0, (0.7 + 0.4)).Transform(transform.Move((0.7+1.13)/2, 0))

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
