// Copyright © 2025 Alex Temnok. All rights reserved.

package holder

import (
	"temnok/pcbc/eda"
	"temnok/pcbc/path"
)

var LIR1254 = &eda.Component{
	Pads: append(
		path.Rect(2.5, 4.5).Clone(2, 15, 0),
		path.Circle(5),
	),

	TrackWidth: 0.25,

	Marks: path.Paths{
		path.Circle(12),
	},

	Components: eda.Components{
		{
			NoClear: true,
			
			Tracks: eda.Tracks(
				eda.Track{{X: -2, Y: -2}, {X: 2, Y: 2}},
				eda.Track{{X: 2, Y: -2}, {X: -2, Y: 2}},
			),
		},
	},
}
