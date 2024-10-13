package sw

import (
	"github.com/stretchr/testify/assert"
	"temnok/pcbc/eda"
	"temnok/pcbc/eda/lib"
	"temnok/pcbc/eda/pcbc"
	"temnok/pcbc/eda/pcbc/bh"
	"temnok/pcbc/eda/pcbc/join"
	"temnok/pcbc/geom"
	"testing"
)

var testBoard = &lib.Component{
	Components: lib.Components{
		pcbc.Board35x45,
		{
			Transform: geom.MoveXY(0, 13).RotateD(-90),
			Components: lib.ComponentsGrid(1, 4, 10.5, 6.5,
				Hyp1TS026A,
				Hyp1TS026A,
				Hyp1TS026A,
				Hyp1TS026A,
			),
		},
		{
			Transform: geom.MoveXY(0, 0).RotateD(0),
			Components: lib.ComponentsGrid(1, 3, 15, 5,
				join.X4,
				join.GndX4,
			),
		},
		{
			Transform: geom.MoveXY(0, -10).RotateD(180),
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
