package pcbc

import (
	"temnok/pcbc/eda/lib"
	"temnok/pcbc/geom"
	"temnok/pcbc/path"
)

var (
	boardHoleContour = path.Paths{path.Circle(2.1)}

	boardHolderHole = &lib.Component{
		Pads:  boardHoleContour,
		Holes: boardHoleContour,
	}

	boardKey = path.Circle(0.6).Transform(geom.MoveXY(-16.25, 21.25))

	Board35x45 = &lib.Component{
		Cuts: path.Paths{path.RoundRect(35, 45, 2.5)},

		Pads: path.Paths{boardKey},

		Components: lib.Components{
			{
				Transform: geom.MoveXY(-15, 20),
				Components: lib.Components{
					boardHolderHole,
				},
			},
			{
				Transform: geom.MoveXY(15, 20),
				Components: lib.Components{
					boardHolderHole,
				},
			},
			{
				Transform: geom.MoveXY(-15, -20),
				Components: lib.Components{
					boardHolderHole,
				},
			},
			{
				Transform: geom.MoveXY(15, -20),
				Components: lib.Components{
					boardHolderHole,
				},
			},
		},
	}
)
