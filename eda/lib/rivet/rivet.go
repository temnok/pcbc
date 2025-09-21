// Copyright © 2025 Alex Temnok. All rights reserved.

package rivet

import (
	"temnok/pcbc/eda"
	"temnok/pcbc/path"
	"temnok/pcbc/util/ptr"
)

var BetweenLayers1and2 = Rivet(1, 2)

func Rivet(layer1, layer2 int) *eda.Component {
	const (
		viaDiameter    = 0.63
		topDiameter    = 1.2
		bottomDiameter = 1.2
	)

	return &eda.Component{
		Inner: eda.Components{
			&eda.Component{
				Layer: layer1,

				Vias: path.Paths{path.Circle(viaDiameter)},

				Tracks: path.Paths{path.Path{path.Point{}}},

				TracksWidth: ptr.To(topDiameter),
			},

			&eda.Component{
				Layer: layer2,

				Tracks: path.Paths{path.Path{path.Point{}}},

				TracksWidth: ptr.To(bottomDiameter),
			},
		},
	}
}
