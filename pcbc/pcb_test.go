package pcbc

import (
	"github.com/stretchr/testify/assert"
	"temnok/lab/eda"
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

	for x := -6.0; x <= 6; x += 12 {
		for y := -10.0; y <= 10; y += 20 {
			pcbc.PY32F002A_QFN16(pcb, geom.MoveXY(x, y).RotateD(-90))
		}
	}

	assert.NoError(t, pcb.SaveFiles("out/"))
}
