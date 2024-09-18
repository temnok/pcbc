package pcbc

import (
	"github.com/stretchr/testify/assert"
	"temnok/lab/contour"
	"temnok/lab/eda"
	"temnok/lab/eda/lib/pcbc"
	"temnok/lab/geom"
	"testing"
)

func TestPCB(t *testing.T) {
	pcb := eda.NewPCB(36, 46)
	pcb.Cut(contour.RoundRect(35, 45, 2.5))

	for y := -9.0; y <= 9; y += 18 {
		pcbc.PY32F002A_QFN16(pcb, geom.MoveXY(0, y))
	}

	assert.NoError(t, pcb.SaveFiles())
}
