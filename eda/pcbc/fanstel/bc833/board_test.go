package bc833

import (
	"github.com/stretchr/testify/assert"
	"temnok/pcbc/eda"
	"temnok/pcbc/eda/lib"
	"temnok/pcbc/eda/pcbc"
	"temnok/pcbc/transform"
	"testing"
)

func Test_BC833(t *testing.T) {
	pcb := eda.NewPCB(36, 46, &lib.Component{
		Components: lib.Components{
			pcbc.Board35x45,
			Board,
		},
	})
	assert.NoError(t, pcb.SaveFiles("gen/long/"))
}

func Test_BC833Short(t *testing.T) {
	pcb := eda.NewPCB(36, 46, &lib.Component{
		Components: lib.Components{
			pcbc.Board35x45,
			ShortBoard.Arrange(transform.Move(0, 10.5)),
			ShortBoard.Arrange(transform.Rotate(180).Move(0, -10.5)),
		},
	})
	assert.NoError(t, pcb.SaveFiles("gen/short/"))
}
