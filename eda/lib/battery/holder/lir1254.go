// Copyright © 2025 Alex Temnok. All rights reserved.

package holder

import (
	"temnok/pcbc/eda"
	"temnok/pcbc/path"
)

var LIR1254 = &eda.Component{
	Pads: path.Join(
		path.Rect(2.5, 4.5).Clone(2, 15, 0),
		path.Circle(5),
	),

	TrackWidth: 0.25,

	GroundTracks: eda.Tracks(
		eda.Track{{X: -2, Y: -2}, {X: 2, Y: 2}},
		eda.Track{{X: 2, Y: -2}, {X: -2, Y: 2}},
	),

	Marks: path.Circle(12),
}
