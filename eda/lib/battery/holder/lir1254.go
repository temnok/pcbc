// Copyright © 2025 Alex Temnok. All rights reserved.

package holder

import (
	"temnok/pcbc/eda"
	"temnok/pcbc/path"
)

var LIR1254 = &eda.Component{
	Pads: append(
		path.RoundRect(2.5, 4.5, 0.5).Clone(2, 15, 0),
		path.Circle(5),
	),

	TracksWidth: 0.5,

	Marks: path.Paths{
		path.Circle(11.5),
	},

	Nested: eda.Components{
		{
			ClearOff: true,

			Tracks: eda.DeprecatedTracks(
				eda.DeprecatedTrack{{X: -2, Y: -2}, {X: 2, Y: 2}},
				eda.DeprecatedTrack{{X: 2, Y: -2}, {X: -2, Y: 2}},
			),
		},
	},
}
