// Copyright © 2025 Alex Temnok. All rights reserved.

package boards

import (
	"temnok/pcbc/eda"
	"temnok/pcbc/path"
	"temnok/pcbc/transform"
)

var (
	hole = &eda.Component{
		CutsFull: true,

		Cuts: path.Paths{path.Circle(2.1)},
	}

	key = path.Paths{path.Circle(0.6).Transform(transform.Move(-16.25, 21.25))}

	Board35x45 = &eda.Component{
		CutsOuter: true,
		Cuts: path.Paths{
			path.RoundRect(35, 45, 2.5),
		},

		Pads: key,

		Nested: eda.Components{
			hole.Clone(2, 30, 0).Clone(2, 0, 40),
		},
	}
)
