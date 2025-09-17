// Copyright © 2025 Alex Temnok. All rights reserved.

package via

import (
	"temnok/pcbc/eda"
	"temnok/pcbc/path"
)

var BetweenLayers1and2 = Via(1, 2)

func Via(layer1, layer2 int) *eda.Component {
	const (
		viaDiameter    = 0.32
		topDiameter    = 1.0
		bottomDiameter = 1.0
	)

	return &eda.Component{
		Inner: eda.Components{
			&eda.Component{
				Layer: layer1,

				Pads: path.Paths{path.Circle(topDiameter)},

				Vias: path.Paths{path.Circle(viaDiameter)},
			},

			&eda.Component{
				Layer: layer2,

				Tracks: path.Paths{path.Path{path.Point{}}},

				TracksWidth: bottomDiameter,
			},
		},
	}
}
