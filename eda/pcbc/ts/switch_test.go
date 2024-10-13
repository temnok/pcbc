package ts

import (
	"github.com/stretchr/testify/assert"
	"temnok/pcbc/eda"
	"temnok/pcbc/eda/lib"
	"temnok/pcbc/eda/pcbc"
	"temnok/pcbc/eda/pcbc/bh/lir1254"
	"temnok/pcbc/eda/pcbc/ts/ts1088"
	"temnok/pcbc/geom"
	"testing"
)

var testBoard = &lib.Component{
	Components: lib.Components{
		pcbc.Board35x45,
		{
			Transform: geom.MoveXY(0, 8),
			Components: lib.Components{
				lir1254.Board,
			},
		},
		{
			Transform: geom.MoveXY(0, -8),
			Components: lib.ComponentsGrid(2, 2, 11, 8,
				ts1088.Board,
				ts1088.Board,
				ts1088.Board,
				ts1088.Board,
			),
		},
	},
}

func TestBoard(t *testing.T) {
	pcb := eda.NewPCB(36, 46)
	pcb.Component(testBoard)

	assert.NoError(t, pcb.SaveFiles("out/"))
}
