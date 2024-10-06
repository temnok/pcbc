package usbc

import (
	"github.com/stretchr/testify/assert"
	"temnok/lab/eda"
	"temnok/lab/eda/lib"
	"temnok/lab/geom"
	"temnok/lab/pcbc"
	"testing"
)

func Test_USBC(t *testing.T) {
	pcb := eda.NewPCB(36, 46)

	pcb.Component(&lib.Component{
		Components: lib.Components{
			pcbc.Board35x45,
			{
				Transform: geom.MoveXY(0, 0),
				Components: lib.Components{
					Board,
				},
			},
		},
	})

	assert.NoError(t, pcb.SaveFiles("out/"))
}
