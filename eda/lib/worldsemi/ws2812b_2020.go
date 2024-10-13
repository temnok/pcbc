package worldsemi

import (
	"temnok/pcbc/eda/lib"
	"temnok/pcbc/geom"
	"temnok/pcbc/path"
)

var (
	pad2020              = path.Rect(0.8, 0.8)
	pad2020_X, pad2020_Y = 0.915, 0.55

	WS2812B_2020 = &lib.Component{
		Pads: path.Paths{
			1: pad2020.Transform(geom.MoveXY(pad2020_X, -pad2020_Y)),
			2: pad2020.Transform(geom.MoveXY(pad2020_X, pad2020_Y)),
			3: pad2020.Transform(geom.MoveXY(-pad2020_X, pad2020_Y)),
			4: pad2020.Transform(geom.MoveXY(-pad2020_X, -pad2020_Y)),
		},
		Marks: path.Strokes{
			0.1: path.Paths{
				path.Rect(2.20, 2),
			},
			0.6: path.Paths{
				path.Path{{0, 0.4}},
			},
			0.2: path.Paths{
				path.Path{{-0.3, -0.4}},
				path.Path{{0, -0.4}},
				path.Path{{0.3, -0.4}},
			},
		},
	}
)
