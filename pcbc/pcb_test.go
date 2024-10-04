package pcbc

import (
	"github.com/stretchr/testify/assert"
	"temnok/lab/eda"
	"temnok/lab/eda/lib"
	"temnok/lab/eda/lib/pcbc"
	"temnok/lab/geom"
	"testing"
)

func TestPCB(t *testing.T) {
	pcb := eda.NewPCB(36, 46)

	pcb.Component(&lib.Component{
		Components: lib.Components{
			pcbc.Board35x45,
			{
				Description: "Copy 1",
				Transform:   geom.MoveXY(0, -12.5),
				Components: lib.Components{
					pcbc.PY32F002A_QFN16(),
				},
			},
			{
				Description: "Copy 2",
				Transform:   geom.MoveXY(0, 0),
				Components: lib.Components{
					pcbc.PY32F002A_QFN16(),
				},
			},
			{
				Description: "Copy 3",
				Transform:   geom.MoveXY(0, 12.5),
				Components: lib.Components{
					pcbc.PY32F002A_QFN16(),
				},
			},
		},
	})

	assert.NoError(t, pcb.SaveFiles("out/"))
}
