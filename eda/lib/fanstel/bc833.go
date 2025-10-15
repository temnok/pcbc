// Copyright © 2025 Alex Temnok. All rights reserved.

package fanstel

import (
	"github.com/temnok/pcbc/eda"
	"github.com/temnok/pcbc/path"
	"github.com/temnok/pcbc/transform"
)

var (
	pad = path.Paths{
		path.Circle(0.6),
	}

	BC833 = &eda.Component{
		Pads: path.Join(
			path.Paths{nil},
			pad.Transform(transform.Move(-4.5, 4.8)),
			pad.Transform(transform.Move(-4.5, 3.8)),
			pad.Transform(transform.Move(-4.5, 2.8)),
			pad.Transform(transform.Move(-4.5, 1.8)),
			pad.Transform(transform.Move(-4.5, 0.8)),
			pad.Transform(transform.Move(-3.5, 3.8)),
			pad.Transform(transform.Move(-3.5, 2.8)),
			pad.Transform(transform.Move(-3.5, 1.8)),
			pad.Transform(transform.Move(-0.5, -0.2)),
			pad.Transform(transform.Move(0.5, -0.2)),
			pad.Transform(transform.Move(1.5, -0.2)),
			pad.Transform(transform.Move(2.5, -0.2)),
			pad.Transform(transform.Move(3.5, -0.2)),
			pad.Transform(transform.Move(4.5, -0.2)),
			pad.Transform(transform.Move(4.5, 1.8)),
			pad.Transform(transform.Move(4.5, 2.8)),
			pad.Transform(transform.Move(4.5, 3.8)),
			pad.Transform(transform.Move(4.5, 4.8)),
			pad.Transform(transform.Move(3.5, 1.8)),
			pad.Transform(transform.Move(3.5, 2.8)),
			pad.Transform(transform.Move(3.5, 3.8)),
			pad.Transform(transform.Move(3.5, 4.8)),
			pad.Transform(transform.Move(2.5, 4.8)),
			pad.Transform(transform.Move(1.5, 4.8)),
			pad.Transform(transform.Move(0.5, 2.8)),
			pad.Transform(transform.Move(0.5, 1.8)),
			pad.Transform(transform.Move(-0.5, 2.8)),
			pad.Transform(transform.Move(-0.5, 1.8)),
		).Transform(transform.Move(0, -(14.3/2 - 1.88))),

		Marks: path.Paths{
			path.Rect(10, 14.3),
			path.Linear([]path.Point{{-5, (14.3 / 2) - 4.4}, {5, (14.3 / 2) - 4.4}}),
		},
	}
)
