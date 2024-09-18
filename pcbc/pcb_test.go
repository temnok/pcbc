package pcbc

import (
	"github.com/stretchr/testify/assert"
	"temnok/lab/contour"
	"temnok/lab/eda"
	"temnok/lab/eda/lib/pkg/pad"
	"temnok/lab/eda/lib/pkg/qfn16"
	"temnok/lab/geom"
	"testing"
)

type XY = geom.XY

func TestPCB(t *testing.T) {
	pcb := eda.NewPCB(36, 46)
	pcb.Cut(geom.Identity(), contour.RoundRect(35, 45, 2.5))

	for y := -9.0; y <= 9; y += 18 {
		py32f002aBoard(geom.Move(XY{0, y}), pcb)
	}

	assert.NoError(t, pcb.SaveFiles())
}

func py32f002aBoard(t geom.Transform, pcb *eda.PCB) {
	textScale := geom.Scale(XY{0.75, 1})
	titleHeight := 2.0
	textHeight := 1.5

	pcb.SilkText(t.MoveXY(-7, 0).Multiply(textScale), titleHeight, "PY32")
	pcb.SilkText(t.MoveXY(2.75, 0).Multiply(textScale), titleHeight, "F002A")

	qfnT := geom.RotateD(45)
	pins := qfnT.Points(qfn16.Add(pcb, t.Multiply(qfnT)))

	n := 9
	padT := geom.MoveXY(0, -6)
	pads := padT.Points(pad.Row(pcb, t.Multiply(padT), contour.Circle(0.75), n, 2.54))
	pad.Row(pcb, t.MoveXY(0, 6), contour.Circle(0.75), n, 2.54)

	for _, t := range []geom.Transform{t, t.RotateD(180)} {
		pcb.Track(t, pads[0], XY{-7.5, -3.2}, XY{-4.5, -3.2}, pins[0])
		pcb.Track(t, pads[1], XY{-6, -4.2}, XY{-4.8, -4.2}, pins[1])
		pcb.Track(t, pads[2], XY{-3.8, -4.5}, XY{-3.8, -4}, pins[2])
		pcb.Track(t, pads[3], XY{pads[3].X, -3.5}, pins[3])
		pcb.Track(t, pads[4], XY{1.5, -4.5}, XY{1.5, -2.5}, pins[4])
		pcb.Track(t, pads[5], XY{pads[5].X, -2.8}, pins[5])
		pcb.Track(t, pads[6], XY{pads[6].X, -4.5}, pins[6])
		pcb.Track(t, pads[7], XY{pads[7].X, -6.2}, pins[7])
		pcb.Track(t, pads[8], XY{pads[8].X, pins[16].Y}, pins[16])
	}

	loPinNames := []string{
		"PB1",
		"PA12",
		"SWD",
		"SWC",
		"PF2",
		"PA0",
		"PA1",
		"PA2",
		"GND",
	}
	hiPinNames := []string{
		"GND",
		"PA8",
		"VCC",
		"PB0",
		"PA7",
		"PA6",
		"PA5",
		"PA4",
		"PA3",
	}
	for i := 0; i < n; i++ {
		pcb.SilkText(t.Move(pads[i]).MoveXY(0, 1.5).RotateD(90).Multiply(textScale), textHeight, loPinNames[i])
		pcb.SilkText(t.Move(pads[i]).MoveXY(0, 8.3).RotateD(90).Multiply(textScale), textHeight, hiPinNames[i])
	}

	pad.Row(pcb, t, contour.Circle(1.3), 2, 20)
	for x := -10.0; x <= 10; x += 20 {
		t := t.Move(XY{x, 0})
		pcb.Pad(t, contour.Circle(1.35))
		pcb.Hole(t, contour.Circle(0.85))
	}

	pcb.Cut(t, contour.RoundRect(24, 16, 1.5))
}
