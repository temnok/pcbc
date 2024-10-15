package qfn16

import (
	"github.com/stretchr/testify/assert"
	"temnok/pcbc/eda"
	"temnok/pcbc/eda/lib"
	"temnok/pcbc/eda/pcbc"
	"temnok/pcbc/eda/pcbc/fanstel/bc833"
	"temnok/pcbc/geom"
	"testing"
)

func TestBoard(t *testing.T) {
	pcb := eda.NewPCB(36, 46)

	pcb.Component(&lib.Component{
		Components: lib.Components{
			pcbc.Board35x45,
			{
				Transform: geom.MoveXY(0, 7),
				Components: lib.Components{
					bc833.ShortBoard,
				},
			},
			{
				Transform: geom.MoveXY(0, -11),
				Components: lib.Components{
					Board,
				},
			},
		},
	})

	assert.NoError(t, pcb.SaveFiles("gen/"))
}
