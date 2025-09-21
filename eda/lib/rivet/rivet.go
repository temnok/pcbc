// Copyright © 2025 Alex Temnok. All rights reserved.

package rivet

import (
	"temnok/pcbc/eda"
	"temnok/pcbc/path"
)

var BetweenLayers1and2 = Rivet(1, 2)

func Rivet(layer1, layer2 int) *eda.Component {
	const (
		viaDiameter    = 0.63
		topDiameter    = 1.2
		bottomDiameter = 1.2
	)

	return &eda.Component{
		Nested: eda.Components{
			{
				Layer: layer1,

				Tracks:      path.Paths{path.Path{path.Point{}}},
				TracksWidth: topDiameter,
			},

			{
				Layer: layer1,

				Cuts:      path.Paths{path.Circle(viaDiameter)},
				CutsInner: true,

				ClearWidth: 0.1,
			},

			{
				Layer: layer2,

				Tracks:      path.Paths{path.Path{path.Point{}}},
				TracksWidth: bottomDiameter,
			},
		},
	}
}
