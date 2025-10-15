// Copyright © 2025 Alex Temnok. All rights reserved.

package rivet

import (
	"github.com/temnok/pcbc/eda"
	"github.com/temnok/pcbc/path"
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
			ClearWidth: 0.1,

			CutsWidth: eda.CutsHidden,
			Cuts:      path.Paths{path.Circle(viaDiameter)},
		},

		{
			Bottom: true,

			Tracks:      path.Paths{path.Path{path.Point{}}},
			TracksWidth: bottomDiameter,
		},
	},
}
