package female

import (
	"github.com/stretchr/testify/assert"
	"temnok/pcbc/eda"
	"temnok/pcbc/eda/lib"
	"testing"
)

func TestBoard(t *testing.T) {
	pcb := eda.NewPCB(2.54*8+5, 10, &lib.Component{
		Components: lib.Components{
			Header254(8),
		},
	})

	assert.NoError(t, pcb.SaveFiles("gen/"))
}
