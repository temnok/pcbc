package bc833

import (
	"github.com/stretchr/testify/assert"
	"temnok/pcbc/eda"
	"temnok/pcbc/eda/lib"
	"temnok/pcbc/eda/pcbc"
	"temnok/pcbc/geom"
	"testing"
)

func Test_BC833(t *testing.T) {
	pcb := eda.NewPCB(36, 46)
	pcb.Component(&lib.Component{
		Components: lib.Components{
			pcbc.Board35x45,
			Board,
		},
	})
	assert.NoError(t, pcb.SaveFiles("gen/long/"))
}

func Test_BC833Short(t *testing.T) {
	pcb := eda.NewPCB(36, 46)
	pcb.Component(&lib.Component{
		Components: lib.Components{
			pcbc.Board35x45,
			ShortBoard.Arrange(geom.MoveXY(0, 10.5)),
			ShortBoard.Arrange(geom.MoveXY(0, -10.5).RotateD(180)),
		},
	})
	assert.NoError(t, pcb.SaveFiles("gen/short/"))
}
