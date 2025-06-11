// Copyright © 2025 Alex Temnok. All rights reserved.

package rivet

import (
	"temnok/pcbc/eda"
	"temnok/pcbc/path"
)

var Rivet06mm_Layers12 = Rivet(0.6, 1.2, 1, 2)

func Rivet(outerDiameter, headDiameter float64, layer1, layer2 int) *eda.Component {
	side := &eda.Component{
		NoOpening: true,

		Pads: path.Paths{path.Circle(headDiameter)},

		Inner: eda.Components{
			{
				ClearWidth: 0.15,

				OuterCut: true,

				Cuts: path.Paths{path.Circle(outerDiameter)},
			},
		},
	}

	return &eda.Component{
		Inner: eda.Components{
			side.WithLayer(layer1),
			side.WithLayer(layer2),
		},
	}
}
