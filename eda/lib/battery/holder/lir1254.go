package holder

import (
	"temnok/pcbc/eda"
	"temnok/pcbc/path"
)

var LIR1254 = &eda.Component{
	Pads: path.Rect(2.5, 4.5).Clone(2, 15, 0).Append(
		path.Paths{path.Circle(5)},
	),

	GroundTracks: path.Strokes{
		0.25: eda.TrackPaths(
			eda.Track{{X: -2, Y: -2}, {X: 2, Y: 2}},
			eda.Track{{X: 2, Y: -2}, {X: -2, Y: 2}},
		),
	},

	Marks: path.Strokes{
		0.1: path.Paths{
			path.Circle(12),
		},
	},
}
