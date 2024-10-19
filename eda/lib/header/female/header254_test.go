package female

import (
	"github.com/stretchr/testify/assert"
	"temnok/pcbc/eda"
	"temnok/pcbc/eda/lib"
	"testing"
)

func TestBoard(t *testing.T) {
	pcb := eda.NewPCB(2.54*40+5, 10)

	pcb.Component(&lib.Component{
		Components: lib.Components{
			Header254(40),
		},
	})

	assert.NoError(t, pcb.SaveFiles("gen/"))
}
