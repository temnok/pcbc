package pcbc

import (
	"temnok/lab/eda"
	"temnok/lab/eda/lib/pkg/qfn16"
	"temnok/lab/geom"
	"testing"
)

func TestPCB(t *testing.T) {
	pcb := eda.NewPCB(25, 25)

	qfn16.Add(pcb, geom.RotateD(45))

	pcb.SaveFiles()
}
