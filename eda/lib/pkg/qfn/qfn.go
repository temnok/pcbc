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

	col := pad.CloneXY(4, 0, -0.5).Transform(transform.Move(-1.6, 0))

	var pads path.Paths
	for a := 0.0; a < 360; a += 90 {
		pads = append(pads, col.Transform(transform.RotateDegrees(a))...)
	}

	//pads = append(pads, path.Rect(1.75, 1.75))

	QFN16G = &eda.Component{
		Pads: pads,

		Marks: path.Paths{
			path.Rect(3.1, 3.1),
			path.Circle(0.2).Transform(transform.Move(-1.7, 1.4)),
		},

		Nested: eda.Components{
			{
				ClearWidth:  eda.ClearOff,
				TracksWidth: 0.25,

				Tracks: path.Paths{
					eda.LinearTrack(path.Point{X: -1.5, Y: 1.5}, path.Point{X: 1.5, Y: -1.5}),
				},
			},

			{
				ClearWidth: 0.3,

				Pads: path.Paths{
					path.Rect(1.75, 1.75),
				},
			},
		},
	}
}
