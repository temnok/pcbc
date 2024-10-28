package qfn

import (
	"temnok/pcbc/eda"
	"temnok/pcbc/path"
	"temnok/pcbc/transform"
)

var QFN16G *eda.Component

func init() {
	pad := path.RoundRect(0.65, 0.25, 0.12)

	col := pad.Clone(4, 0, -0.5).Apply(transform.Move(-1.55, 0))

	pads := path.Paths{}
	for a := 0.0; a < 360; a += 90 {
		pads = append(pads, col.Apply(transform.Rotate(a))...)
	}

	pads = append(pads, path.Rect(1.2, 1.2))

	QFN16G = &eda.Component{
		Clears: path.Paths{
			path.Rect(2.2, 2.2),
		},

		Pads: pads,

		Marks: path.Strokes{
			0.1: path.Paths{
				path.Rect(3.1, 3.1),
				path.Circle(0.3).Apply(transform.Move(-1.8, 1.4)),
			},
		},

		GroundTracks: path.Strokes{
			0.16: eda.TrackPaths(
				eda.Track{{X: -1.5, Y: 1.5}, {X: 1.5, Y: -1.5}},
			),
		},
	}
}
