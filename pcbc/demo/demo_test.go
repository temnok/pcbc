package demo

import (
	"github.com/stretchr/testify/assert"
	"temnok/lab/eda"
	"temnok/lab/eda/lib"
	"temnok/lab/geom"
	"temnok/lab/pcbc"
	"temnok/lab/pcbc/py32f002a"
	"temnok/lab/pcbc/usbc"
	"testing"
)

func Test_Demo(t *testing.T) {
	pcb := eda.NewPCB(36, 46)

	pcb.Component(&lib.Component{
		Components: lib.Components{
			pcbc.Board35x45,
			{
				Transform: geom.MoveXY(0, 9),
				Components: lib.Components{
					py32f002a.QFN16,
				},
			},
			{
				Transform: geom.MoveXY(-6.5, -7),
				Components: lib.Components{
					py32f002a.SOP8,
				},
			},
			{
				Transform: geom.MoveXY(6.5, -7).RotateD(90),
				Components: lib.Components{
					usbc.Board,
				},
			},
		},
	})

	assert.NoError(t, pcb.SaveFiles("out/"))
}
