// Copyright © 2025 Alex Temnok. All rights reserved.

package boards

import (
	"temnok/pcbc/eda"
	"temnok/pcbc/path"
	"temnok/pcbc/transform"
)

var MountHole = &eda.Component{
	Cuts: path.Paths{path.Circle(1.5)},

	TracksWidth: 0.5,

	Nested: eda.Components{
		{
			Pads: path.Pie(6, 1.0, 1.3, 20).Transform(transform.RotateDegrees(-30)),

			Nested: eda.Components{
				{
					Tracks: path.Paths{path.Circle(2.3)},
				},

				{
					ClearOff: true,

					Tracks: path.Paths{
						eda.Track(path.Point{X: 1.25, Y: 0}, path.Point{X: 1.25 + 0.35, Y: 0}),
						eda.Track(path.Point{X: -1.25, Y: 0}, path.Point{X: -1.25 - 0.35, Y: 0}),
						eda.Track(path.Point{X: 0, Y: 1.25}, path.Point{X: 0, Y: 1.25 + 0.35}),
						eda.Track(path.Point{X: 0, Y: -1.25}, path.Point{X: 0, Y: -1.25 - 0.35}),
					},
				},
			},
		},
	},
}
