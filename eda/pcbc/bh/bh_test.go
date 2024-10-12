package bh

import (
	"github.com/stretchr/testify/assert"
	"temnok/pcbc/eda"
	"temnok/pcbc/eda/lib"
	"temnok/pcbc/eda/pcbc"
	"testing"
)

var testBoard = &lib.Component{
	//Clears: path.Paths{path.Rect(36, 46)},
	Components: lib.Components{
		pcbc.Board35x45,
		{
			//Transform: geom.RotateD(-90),
			Components: lib.ComponentsGrid(1, 1, 10.5, 5.5,
				LIR1254,
			),
		},
	},
}

func TestBatteryHolder(t *testing.T) {
	pcb := eda.NewPCB(36, 46)
	pcb.Component(testBoard)
	assert.NoError(t, pcb.SaveFiles("out/"))
}
