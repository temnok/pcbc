// Copyright © 2025 Alex Temnok. All rights reserved.

package boards

import (
	"github.com/temnok/pcbc/eda"
	"github.com/temnok/pcbc/path"
	"github.com/temnok/pcbc/transform"
)

var MountHole = &eda.Component{
	Cuts: path.Paths{path.Circle(1.43)},

	Nested: eda.Components{
		alignMark,

		(&eda.Component{
			ClearWidth: eda.ClearOff,
			Pads: path.Paths{
				path.RoundRect(0.3, 0.7, 0.15).Transform(transform.Move(1.1, 0)),
			},
		}).Clone(6, transform.RotateDegrees(60)),
	},
}
