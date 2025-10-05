// Copyright © 2025 Alex Temnok. All rights reserved.

package boards

import (
	"temnok/pcbc/eda"
	"temnok/pcbc/path"
	"temnok/pcbc/transform"
)

var MountHoleV2 = &eda.Component{
	Cuts: path.Paths{path.Circle(1.43)},

	Nested: eda.Components{
		(&eda.Component{
			ClearWidth: eda.ClearOff,
			//Pads:       path.Pie(6, 1.0, 1.3, 20),
			Pads: path.Paths{
				path.RoundRect(0.3, 0.7, 0.15).Transform(transform.Move(1.1, 0)),
			},
		}).Clone(6, transform.RotateDegrees(60)),
	},
}
