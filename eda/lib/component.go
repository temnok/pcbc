package lib

import (
	"temnok/lab/geom"
	"temnok/lab/path"
)

type Component struct {
	Pads path.Paths
}

func (c *Component) Transform(t geom.Transform) *Component {
	return &Component{
		Pads: c.Pads.Transform(t),
	}
}
