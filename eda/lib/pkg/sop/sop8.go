// Copyright © 2025 Alex Temnok. All rights reserved.

package sop

import (
	"github.com/temnok/pcbc/eda"
	"github.com/temnok/pcbc/path"
	"github.com/temnok/pcbc/transform"
)

var (
	pad = path.RoundRect(0.5, 1.0, 0.25)

	bottomPads = pad.CloneXY(4, 1.27, 0).Transform(transform.Move(0, -2.75))

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
