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

	for y := -12.5; y <= 12.5; y += 12.5 {
		c := pcbc.PY32F002A_QFN16()
		pcb.Component(c.Transform(geom.MoveXY(0, y)))
	}

	assert.NoError(t, pcb.SaveFiles("out/"))
}
