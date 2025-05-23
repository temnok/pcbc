// Copyright © 2025 Alex Temnok. All rights reserved.

package qfn

import (
	"temnok/pcbc/eda"
	"temnok/pcbc/path"
	"temnok/pcbc/transform"
)

var QFN16G *eda.Component

func init() {
	pad := path.RoundRect(0.7, 0.24, 0.1)

	col := pad.Clone(4, 0, -0.5).Transform(transform.Move(-1.6, 0))

	var pads path.Paths
	for a := 0.0; a < 360; a += 90 {
		pads = append(pads, col.Transform(transform.RotateDegrees(a))...)
	}

	pads = append(pads, path.Rect(1.5, 1.5))

	QFN16G = &eda.Component{
		//Clears: path.Paths{path.Rect(2.5, 2.5)},

		Pads: pads,

		Marks: path.Paths{
			path.Rect(3.1, 3.1),
			path.Circle(0.2).Transform(transform.Move(-1.7, 1.4)),
		},

		TrackWidth: 0.25,

		GroundTracks: eda.Tracks(
			eda.Track{{X: -1.5, Y: 1.5}, {X: 1.5, Y: -1.5}},
		),
	}
}
