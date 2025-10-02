// Copyright © 2025 Alex Temnok. All rights reserved.

package holder

import (
	"temnok/pcbc/eda"
	"temnok/pcbc/path"
)

var LIR1254 = &eda.Component{
	Pads: path.RoundRect(1.75, 4.5, 0.5).CloneXY(2, 14, 0),

	Marks: path.Paths{
		path.Circle(11.5),
	},

	Nested: eda.Components{
		{
			ClearWidth: eda.ClearOff,

			Pads: path.Paths{
				path.Circle(4.5),
			},
		},
	},
}
