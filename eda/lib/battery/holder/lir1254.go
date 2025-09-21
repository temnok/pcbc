// Copyright © 2025 Alex Temnok. All rights reserved.

package holder

import (
	"temnok/pcbc/eda"
	"temnok/pcbc/path"
	"temnok/pcbc/util/ptr"
)

var LIR1254 = &eda.Component{
	Pads: append(
		path.RoundRect(2.5, 4.5, 0.5).Clone(2, 15, 0),
		path.Circle(5),
	),

	TracksWidth: ptr.To(0.5),

	Marks: path.Paths{
		path.Circle(12),
	},

	Inner: eda.Components{
		{
			ClearNone: true,

			Tracks: eda.Tracks(
				eda.Track{{X: -2, Y: -2}, {X: 2, Y: 2}},
				eda.Track{{X: 2, Y: -2}, {X: -2, Y: 2}},
			),
		},
	},
}
