package pcbc

import (
	"temnok/pcbc/eda"
	"temnok/pcbc/path"
	"temnok/pcbc/transform"
)

var (
	MountHole = &eda.Component{
		Pads: path.Pie(6, 1.0, 1.3, 15).Apply(transform.Rotate(-30)),

		GroundTracks: eda.TrackPaths(
			eda.Track{{1.15, 0}}.DX(0.6),
			eda.Track{{-1.15, 0}}.DX(-0.6),
		),

		Openings: path.Paths{
			path.Circle(2.6),
		},

		Holes: path.Paths{
			path.Circle(1.8),
		},

		Components: eda.Components{
			{
				TrackThickness: 0.2,
				Tracks:         path.Paths{path.Circle(2.8)},
			},
		},
	}
)
