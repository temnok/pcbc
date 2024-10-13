package fanstel

import (
	"github.com/stretchr/testify/assert"
	"temnok/pcbc/eda"
	"temnok/pcbc/eda/lib"
	"temnok/pcbc/geom"
	"temnok/pcbc/path"
	"testing"
)

func TestBoard(t *testing.T) {
	pcb := eda.NewPCB(20, 20)

	x, y := 5.0, 14.3/2

	pcb.Component(&lib.Component{
		Pads: path.Paths{
			path.Circle(0.5).Transform(geom.MoveXY(x, y)),
			path.Circle(0.5).Transform(geom.MoveXY(x, -y)),
			path.Circle(0.5).Transform(geom.MoveXY(-x, -y)),
			path.Circle(0.5).Transform(geom.MoveXY(-x, y)),
		},
		Components: lib.Components{
			BC833,
		},
	})

	assert.NoError(t, pcb.SaveFiles("out/"))
}
