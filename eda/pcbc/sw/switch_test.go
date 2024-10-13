package sw

import (
	"github.com/stretchr/testify/assert"
	"temnok/pcbc/eda"
	"temnok/pcbc/eda/lib"
	"temnok/pcbc/eda/pcbc"
	"temnok/pcbc/eda/pcbc/bh"
	"temnok/pcbc/geom"
	"testing"
)

var testBoard = &lib.Component{
	Components: lib.Components{
		pcbc.Board35x45,
		{
			Transform: geom.MoveXY(0, 8),
			Components: lib.ComponentsGrid(2, 2, 11, 8,
				XunpuTS1088,
				XunpuTS1088,
				XunpuTS1088,
				XunpuTS1088,
			),
		},
		{
			Transform: geom.MoveXY(0, -8).RotateD(180),
			Components: lib.Components{
				bh.LIR1254,
			},
		},
	},
}

func TestSwitch(t *testing.T) {
	pcb := eda.NewPCB(36, 46)
	pcb.Component(testBoard)

	assert.NoError(t, pcb.SaveFiles("out/"))
}
