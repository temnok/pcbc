package pcbc

import (
	"temnok/pcbc/eda"
	"temnok/pcbc/path"
	"temnok/pcbc/transform"
)

var MountHole = &eda.Component{
	Tracks: path.Strokes{
		0.2: {
			path.Circle(2.7),
		},
	},

	GroundTracks: path.Strokes{
		0.25: {
			path.Lines(path.Points{{1.7, 0}, {-1.7, 0}}),
		},
	},

	Openings: path.Paths{
		path.Circle(2.6),
	},

	Pads: path.Pie(6, 1.0, 1.3, 15).Apply(transform.Rotate(-30)),

	Holes: path.Paths{
		path.Circle(1.8),
	},
}
