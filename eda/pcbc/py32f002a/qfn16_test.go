package py32f002a

import (
	"github.com/stretchr/testify/assert"
	"temnok/pcbc/eda"
	"temnok/pcbc/eda/lib"
	"temnok/pcbc/eda/pcbc"
	"temnok/pcbc/geom"
	"testing"
)

func Test_QFN16(t *testing.T) {
	pcb := eda.NewPCB(36, 46)

	pcb.Component(&lib.Component{
		Components: lib.Components{
			pcbc.Board35x45,
			{
				Description: "Copy 1",
				Transform:   geom.MoveXY(0, -12.5),
				Components: lib.Components{
					QFN16,
				},
			},
			{
				Description: "Copy 2",
				Transform:   geom.MoveXY(0, 0),
				Components: lib.Components{
					QFN16,
				},
			},
			{
				Description: "Copy 3",
				Transform:   geom.MoveXY(0, 12.5),
				Components: lib.Components{
					QFN16,
				},
			},
		},
	})

	assert.NoError(t, pcb.SaveFiles("out/qfn16/"))
}
