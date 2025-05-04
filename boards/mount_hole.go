// Copyright © 2025 Alex Temnok. All rights reserved.

package boards

import (
	"temnok/pcbc/eda"
	"temnok/pcbc/path"
	"temnok/pcbc/transform"
)

var (
	MountHole = &eda.Component{
		Pads: path.Pie(6, 1.0, 1.3, 15).Apply(transform.RotateDegrees(-30)),

		GroundTracks: eda.Tracks(
			eda.Track{{X: 1.15, Y: 0}}.DX(0.6),
			eda.Track{{X: -1.15, Y: 0}}.DX(-0.6),
		),

		Openings: path.Paths{path.Circle(2.6)},

		Holes: path.Paths{path.Circle(1.8)},

		Components: eda.Components{
			{
				TrackWidth: 0.15,
				Tracks:     path.Paths{path.Circle(2.8)},
			},
		},
	}
)
