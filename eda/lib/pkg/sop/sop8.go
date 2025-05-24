// Copyright © 2025 Alex Temnok. All rights reserved.

package sop

import (
	"temnok/pcbc/eda"
	"temnok/pcbc/path"
	"temnok/pcbc/transform"
)

var (
	pad = path.RoundRect(0.5, 1.0, 0.25)

	bottomPads = pad.Clone(4, 1.27, 0).Transform(transform.Move(0, -2.75))

	topPads = bottomPads.Transform(transform.RotateDegrees(180))

	pads = append(bottomPads, topPads...)

	SOP8 = &eda.Component{
		Pads: pads,

		Marks: path.Paths{
			path.Rect(5, 4),
			path.Circle(0.3).Transform(transform.Move(-2.4, -2.3)),
		},
	}
)
