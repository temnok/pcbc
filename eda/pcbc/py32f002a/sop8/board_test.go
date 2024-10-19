package sop8

import (
	"github.com/stretchr/testify/assert"
	"temnok/pcbc/eda"
	"temnok/pcbc/eda/lib"
	"temnok/pcbc/eda/pcbc"
	"temnok/pcbc/eda/pcbc/py32f002a/qfn16"
	"temnok/pcbc/transform"
	"testing"
)

func TestBoard(t *testing.T) {
	pcb := eda.NewPCB(36, 46, &lib.Component{
		Components: lib.Components{
			pcbc.Board35x45,
			qfn16.Board.Arrange(transform.Move(0, 9)),
			Board.Arrange(transform.Move(-6.5, -7)),
			Board.Arrange(transform.Move(6.5, -7)),
		},
	})

	assert.NoError(t, pcb.SaveFiles("gen/"))
}
