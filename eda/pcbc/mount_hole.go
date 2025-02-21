// Copyright © 2025 Alex Temnok. All rights reserved.

package pcbc

import (
	"temnok/pcbc/eda"
	"temnok/pcbc/path"
	"temnok/pcbc/transform"
)

var (
	MountHole = &eda.Component{
		Pads: path.Pie(6, 1.0, 1.3, 15).Apply(transform.Rotate(-30)),

		GroundTracks: eda.Tracks(
			eda.Track{{X: 1.15, Y: 0}}.DX(0.6),
			eda.Track{{X: -1.15, Y: 0}}.DX(-0.6),
		),

		Openings: path.Circle(2.6),

		Holes: path.Circle(1.8),

		Components: eda.Components{
			{
				TrackWidth: 0.2,
				Tracks:     path.Circle(2.8),
			},
		},
	}
)
