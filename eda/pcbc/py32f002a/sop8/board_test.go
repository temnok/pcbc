package sop8

import (
	"github.com/stretchr/testify/assert"
	"temnok/pcbc/eda"
	"temnok/pcbc/eda/lib"
	"temnok/pcbc/eda/pcbc"
	"temnok/pcbc/eda/pcbc/py32f002a/qfn16"
	"temnok/pcbc/geom"
	"testing"
)

func TestBoard(t *testing.T) {
	pcb := eda.NewPCB(36, 46)

	pcb.Component(&lib.Component{
		Components: lib.Components{
			pcbc.Board35x45,
			{
				Transform: geom.MoveXY(0, 9),
				Components: lib.Components{
					qfn16.Board,
				},
			},
			{
				Transform: geom.MoveXY(-6.5, -7),
				Components: lib.Components{
					Board,
				},
			},
			{
				Transform: geom.MoveXY(6.5, -7),
				Components: lib.Components{
					Board,
				},
			},
		},
	})

	assert.NoError(t, pcb.SaveFiles("gen/sop8/"))
}
