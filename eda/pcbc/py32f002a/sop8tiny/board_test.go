package sop8tiny

import (
	"github.com/stretchr/testify/assert"
	"temnok/pcbc/eda"
	"temnok/pcbc/eda/lib"
	"temnok/pcbc/eda/pcbc"
	"temnok/pcbc/transform"
	"testing"
)

func TestBoard(t *testing.T) {
	pcb := eda.NewPCB(36, 46)

	pcb.Component(&lib.Component{
		Components: lib.Components{
			pcbc.Board35x45,

			Board.Arrange(transform.Rotate(-90)).Clone(3, 9.5, 0).Clone(2, 0, 17.5),
		},
	})

	assert.NoError(t, pcb.SaveFiles("gen/"))
}
