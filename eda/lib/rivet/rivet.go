// Copyright © 2025 Alex Temnok. All rights reserved.

package rivet

import (
	"temnok/pcbc/eda"
	"temnok/pcbc/path"
)

const (
	viaDiameter    = 0.63
	topDiameter    = 1.2
	bottomDiameter = 1.2
)

var Rivet = &eda.Component{
	Nested: eda.Components{
		{
			Tracks:      path.Paths{path.Path{path.Point{}}},
			TracksWidth: topDiameter,
		},

		{
			Cuts:       path.Paths{path.Circle(viaDiameter)},
			CutsHidden: true,

			ClearWidth: 0.1,
		},

		{
			Back: true,

			Tracks:      path.Paths{path.Path{path.Point{}}},
			TracksWidth: bottomDiameter,
		},
	},
}
