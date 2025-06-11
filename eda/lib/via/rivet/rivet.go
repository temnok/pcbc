// Copyright © 2025 Alex Temnok. All rights reserved.

package rivet

import (
	"temnok/pcbc/eda"
	"temnok/pcbc/path"
)

var Rivet06mm = &eda.Component{
	NoOpening: true,

	Pads: path.Paths{path.Circle(1.2)},

	Components: eda.Components{
		{
			ClearWidth: 0.15,

			OuterCut: true,

			Cuts: path.Paths{path.Circle(0.6)},
		},
	},
}
