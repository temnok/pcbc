package pcbc

import (
	"temnok/pcbc/eda/lib"
	"temnok/pcbc/path"
	"temnok/pcbc/transform"
)

var (
	boardHoleContour = path.Paths{path.Circle(2.1)}

	boardHolderHole = &lib.Component{
		Pads:  boardHoleContour,
		Holes: boardHoleContour,
	}

	boardKey = path.Circle(0.6).Apply(transform.Move(-16.25, 21.25))

	Board35x45 = &lib.Component{
		Cuts: path.Paths{path.RoundRect(35, 45, 2.5)},

		Pads: path.Paths{boardKey},

		Components: lib.Components{
			boardHolderHole.Arrange(transform.Move(-15, 20)),
			boardHolderHole.Arrange(transform.Move(15, 20)),
			boardHolderHole.Arrange(transform.Move(-15, -20)),
			boardHolderHole.Arrange(transform.Move(15, -20)),
		},
	}
)
