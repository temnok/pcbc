// Copyright © 2025 Alex Temnok. All rights reserved.

package fanstel

import (
	"temnok/pcbc/eda"
	"temnok/pcbc/path"
	"temnok/pcbc/transform"
)

var (
	pad = path.Paths{
		path.Circle(0.6),
	}

	BC833 = &eda.Component{
		Pads: path.Join(
			path.Paths{nil},
			pad.Apply(transform.Move(-4.5, 4.8)),
			pad.Apply(transform.Move(-4.5, 3.8)),
			pad.Apply(transform.Move(-4.5, 2.8)),
			pad.Apply(transform.Move(-4.5, 1.8)),
			pad.Apply(transform.Move(-4.5, 0.8)),
			pad.Apply(transform.Move(-3.5, 3.8)),
			pad.Apply(transform.Move(-3.5, 2.8)),
			pad.Apply(transform.Move(-3.5, 1.8)),
			pad.Apply(transform.Move(-0.5, -0.2)),
			pad.Apply(transform.Move(0.5, -0.2)),
			pad.Apply(transform.Move(1.5, -0.2)),
			pad.Apply(transform.Move(2.5, -0.2)),
			pad.Apply(transform.Move(3.5, -0.2)),
			pad.Apply(transform.Move(4.5, -0.2)),
			pad.Apply(transform.Move(4.5, 1.8)),
			pad.Apply(transform.Move(4.5, 2.8)),
			pad.Apply(transform.Move(4.5, 3.8)),
			pad.Apply(transform.Move(4.5, 4.8)),
			pad.Apply(transform.Move(3.5, 1.8)),
			pad.Apply(transform.Move(3.5, 2.8)),
			pad.Apply(transform.Move(3.5, 3.8)),
			pad.Apply(transform.Move(3.5, 4.8)),
			pad.Apply(transform.Move(2.5, 4.8)),
			pad.Apply(transform.Move(1.5, 4.8)),
			pad.Apply(transform.Move(0.5, 2.8)),
			pad.Apply(transform.Move(0.5, 1.8)),
			pad.Apply(transform.Move(-0.5, 2.8)),
			pad.Apply(transform.Move(-0.5, 1.8)),
		).Apply(transform.Move(0.15, 1.75-14.3/2)),

		Marks: path.Paths{path.Rect(10, 14.3)},
	}
)
