package qfn16tiny

import (
	"github.com/stretchr/testify/assert"
	"temnok/pcbc/eda"
	"temnok/pcbc/eda/lib"
	"temnok/pcbc/eda/pcbc"
	"testing"
)

func TestBoard(t *testing.T) {
	pcb := eda.NewPCB(36, 46, &lib.Component{
		Components: lib.Components{
			pcbc.Board35x45,

			Board,
		},
	})

	assert.NoError(t, pcb.SaveFiles("gen/"))
}
