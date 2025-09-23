// Copyright © 2025 Alex Temnok. All rights reserved.

package boards

import (
	"temnok/pcbc/eda"
	"temnok/pcbc/path"
	"temnok/pcbc/transform"
)

var MountHole15 = &eda.Component{
	Cuts: path.Paths{path.Circle(1.5)},

	TracksWidth: 0.4,

	Nested: eda.Components{
		{
			Pads: path.Pie(6, 1.0, 1.3, 20).Transform(transform.RotateDegrees(-30)),

			Nested: eda.Components{
				{
					Tracks: path.Paths{path.Circle(2.3)},
				},

				{
					ClearOff: true,

					Tracks: eda.Tracks(
						eda.Track{{X: 1.25, Y: 0}}.DX(0.3),
						eda.Track{{X: -1.25, Y: 0}}.DX(-0.3),
						eda.Track{{X: 0, Y: 1.25}}.DY(0.3),
						eda.Track{{X: 0, Y: -1.25}}.DY(-0.3),
					),
				},
			},
		},
	},
}
