package pcbc

import (
	"temnok/pcbc/eda"
	"temnok/pcbc/path"
	"temnok/pcbc/transform"
)

var (
	mountPads = path.Pie(6, 1.0, 1.3, 15).Apply(transform.Rotate(-30))
	pad       = mountPads.Centers(transform.Identity)

	MountHole = &eda.Component{
		Pads: mountPads,

		GroundTracks: eda.TrackPaths(
			eda.Track{pad[0]}.DX(0.7),
			eda.Track{pad[3]}.DX(-0.7),
		),

		Openings: path.Paths{
			path.Circle(2.6),
		},

		Holes: path.Paths{
			path.Circle(1.8),
		},

		Components: eda.Components{
			{
				TrackWidth: 0.2,
				Tracks:     path.Paths{path.Circle(2.8)},
			},
		},
	}
)
