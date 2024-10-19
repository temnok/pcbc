package female

import (
	"github.com/stretchr/testify/assert"
	"temnok/pcbc/eda"
	"temnok/pcbc/eda/lib"
	"testing"
)

func TestBoard(t *testing.T) {
	pcb := eda.NewPCB(30, 20)

	pcb.Component(&lib.Component{
		Components: lib.Components{
			Header254(8),
		},
	})

	assert.NoError(t, pcb.SaveFiles("gen/"))
}
