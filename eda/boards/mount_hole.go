// Copyright © 2025 Alex Temnok. All rights reserved.

package boards

import (
	"temnok/pcbc/eda"
	"temnok/pcbc/path"
	"temnok/pcbc/transform"
	"temnok/pcbc/util/ptr"
)

var (
	MountHole = &eda.Component{
		Pads: path.Pie(6, 1.05, 1.4, 20).Transform(transform.RotateDegrees(-30)),

		Cuts: path.Paths{path.Circle(1.75)},

		Inner: eda.Components{
			{
				TracksWidth: ptr.To(0.35),
				Tracks:      path.Paths{path.Circle(2.45)},
			},

			{
				ClearWidth: ptr.To(0.0),

				Tracks: eda.Tracks(
					eda.Track{{X: 1.15, Y: 0}}.DX(0.5),
					eda.Track{{X: -1.15, Y: 0}}.DX(-0.5),
					eda.Track{{X: 0, Y: 1.15}}.DY(0.5),
					eda.Track{{X: 0, Y: -1.15}}.DY(-0.5),
				),
			},
		},
	}
)
