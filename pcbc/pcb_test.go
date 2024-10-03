package pcbc

import (
	"github.com/stretchr/testify/assert"
	"temnok/lab/eda"
	"temnok/lab/eda/lib"
	"temnok/lab/eda/lib/pcbc"
	"temnok/lab/geom"
	"temnok/lab/path"
	"testing"
)

func TestPCB(t *testing.T) {
	pcb := eda.NewPCB(36, 46)
	board := path.RoundRect(35, 45, 2.5)
	pcb.Cut(board)
	pcb.StencilCut(board)

	pcb.Component(&lib.Component{
		Components: lib.Components{
			{
				Description:    "Copy 1",
				Transformation: geom.MoveXY(0, -12.5),
				Components: lib.Components{
					pcbc.PY32F002A_QFN16(),
				},
			},
			{
				Description:    "Copy 2",
				Transformation: geom.MoveXY(0, 0),
				Components: lib.Components{
					pcbc.PY32F002A_QFN16(),
				},
			},
			{
				Description:    "Copy 3",
				Transformation: geom.MoveXY(0, 12.5),
				Components: lib.Components{
					pcbc.PY32F002A_QFN16(),
				},
			},
		},
	})

	assert.NoError(t, pcb.SaveFiles("out/"))
}
