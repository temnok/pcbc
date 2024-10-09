package pcbc

import (
	"temnok/pcbc/eda/lib"
	"temnok/pcbc/geom"
	"temnok/pcbc/path"
)

var MountHole = &lib.Component{
	Description: "PCBC mount hole",
	Tracks:      path.Strokes{0.2: {path.Circle(2.7)}},
	Openings:    path.Paths{path.Circle(2.6)},
	Pads:        path.Pie(6, 1.0, 1.3, 15*geom.Degree).Transform(geom.RotateD(-30)),
	Holes:       path.Paths{path.Circle(1.8)},
}
