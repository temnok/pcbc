package py32f002a

import (
	"github.com/stretchr/testify/assert"
	"temnok/lab/eda"
	"temnok/lab/eda/lib"
	"temnok/lab/geom"
	"temnok/lab/pcbc"
	"testing"
)

func Test_SOP8(t *testing.T) {
	pcb := eda.NewPCB(36, 46)

	pcb.Component(&lib.Component{
		Components: lib.Components{
			pcbc.Board35x45,
			{
				Transform: geom.MoveXY(0, 9),
				Components: lib.Components{
					QFN16,
				},
			},
			{
				Transform: geom.MoveXY(-6.5, -7),
				Components: lib.Components{
					SOP8,
				},
			},
			{
				Transform: geom.MoveXY(6.5, -7),
				Components: lib.Components{
					SOP8,
				},
			},
		},
	})

	assert.NoError(t, pcb.SaveFiles("out/sop8/"))
}
