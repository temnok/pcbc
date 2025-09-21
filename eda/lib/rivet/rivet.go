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
			&eda.Component{
				Layer: layer1,

				Cuts:     path.Paths{path.Circle(viaDiameter)},
				CutsVias: true,

				Tracks:      path.Paths{path.Path{path.Point{}}},
				TracksWidth: topDiameter,
			},

			&eda.Component{
				Layer: layer2,

				Tracks:      path.Paths{path.Path{path.Point{}}},
				TracksWidth: bottomDiameter,
			},
		},
	}
}
