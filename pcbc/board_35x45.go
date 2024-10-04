package pcbc

import (
	"temnok/lab/eda/lib"
	"temnok/lab/geom"
	"temnok/lab/path"
)

var (
	boardHoleContour = path.Paths{path.Circle(2.1)}

	boardHolderHole = &lib.Component{
		Pads:          boardHoleContour,
		Holes:         boardHoleContour,
		Openings:      boardHoleContour,
		MaskBaseHoles: path.Paths{path.Circle(1.3)},
	}

	boardKey = path.Paths{path.Lines(path.Points{
		{0.4, -0.4},
		{0.3, 0.4},
		{-0.4, -0.3},
		{0.4, -0.4},
	})}.Transform(geom.MoveXY(-16.4, 21.4))

	Board35x45 = &lib.Component{
		Cuts: path.Paths{path.RoundRect(35, 45, 2.5)},

		Pads: boardKey,
		Marks: path.Strokes{
			0.2: boardKey,
		},

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
