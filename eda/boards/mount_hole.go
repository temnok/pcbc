// Copyright © 2025 Alex Temnok. All rights reserved.

package boards

import (
	"temnok/pcbc/eda"
	"temnok/pcbc/path"
	"temnok/pcbc/transform"
)

var (
	MountHole = &eda.Component{
		Pads: path.Pie(6, 1.1, 1.35, 15).Transform(transform.RotateDegrees(-30)),

		Cuts: path.Paths{path.Circle(1.8)},

		Components: eda.Components{
			{
				TrackWidth: 0.28,
				Tracks:     path.Paths{path.Circle(2.45)},
			},

			{
				NoClear: true,

				Tracks: eda.Tracks(
					eda.Track{{X: 1.2, Y: 0}}.DX(0.5),
					eda.Track{{X: -1.2, Y: 0}}.DX(-0.5),
				),
			},
		},
	}
)
