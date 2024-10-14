package ws2812b_2020

import (
	"github.com/stretchr/testify/assert"
	"temnok/pcbc/eda"
	"temnok/pcbc/eda/lib"
	"temnok/pcbc/eda/pcbc"
	"temnok/pcbc/eda/pcbc/fanstel/bc833"
	"temnok/pcbc/geom"
	"testing"
)

var testBoard = &lib.Component{
	Components: lib.Components{
		pcbc.Board35x45,
		{
			Transform: geom.MoveXY(0, 3.5),
			Components: lib.Components{
				bc833.Board,
			},
		},
		{
			Transform: geom.MoveXY(0, -13),
			Components: lib.ComponentsGrid(2, 1, 11, 7,
				Board,
				Board,
				Board,
				Board,
			),
		},
	},
}

func TestBoard(t *testing.T) {
	pcb := eda.NewPCB(36, 46)
	pcb.Component(testBoard)

	assert.NoError(t, pcb.SaveFiles("gen/"))
}
