// Copyright © 2025 Alex Temnok. All rights reserved.

package rivet

import (
	"temnok/pcbc/eda"
	"temnok/pcbc/path"
)

var Rivet06mm = &eda.Component{
	NoOpening: true,

	TrackWidth: 0.2,

	Pads: path.Paths{path.Circle(1)},

	Components: eda.Components{
		{
			OuterCut: true,

			Cuts: path.Paths{path.Circle(0.6)},
		},
	},
}
