package pcbc

import (
	"temnok/lab/eda/lib"
	"temnok/lab/geom"
	"temnok/lab/path"
)

var MountHole = &lib.Component{
	Description: "PCBC mount hole",
	Tracks:      path.Strokes{0.2: {path.Circle(2.7)}},
	Openings:    path.Paths{path.Circle(2.6)},
	Pads:        path.Pie(6, 1.0, 1.3, 10*geom.Degree),
	Holes:       path.Paths{path.Circle(1.8)},
}
