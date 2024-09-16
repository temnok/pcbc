package pcbc

import (
	"temnok/lab/contour"
	"temnok/lab/eda"
	"temnok/lab/eda/lib/pkg/pad"
	"temnok/lab/eda/lib/pkg/qfn16"
	"temnok/lab/geom"
	"testing"
)

type XY = geom.XY

func TestPCB(t *testing.T) {
	pcb := eda.NewPCB(38, 48)
	for y := -9.0; y <= 9; y += 18 {
		pcb.With(func() {
			pcb.Transform = pcb.Transform.Multiply(geom.MoveXY(0, y))

			qfn16pinout(t, pcb)
		})
	}

	pcb.SaveFiles()
}

func qfn16pinout(t *testing.T, pcb *eda.PCB) {
	textScale := geom.Scale(XY{0.75, 1})
	textHeight := 1.5

	pcb.SilkText(geom.MoveXY(-5.5, 0.75).Multiply(textScale), textHeight, "PY32")
	pcb.SilkText(geom.MoveXY(-6.25, -0.75).Multiply(textScale), textHeight, "F002A")
	pcb.SilkText(geom.MoveXY(2.75, 0.75).Multiply(textScale), textHeight, "W15U")
	pcb.SilkText(geom.MoveXY(2.75, -0.75).Multiply(textScale), textHeight, "6TR")

	pins := qfn16.Add(pcb, geom.RotateD(45))

	n := 9
	pads := pad.Row(pcb, geom.MoveXY(0, -6), contour.Circle(0.75), n, 2.54)
	pad.Row(pcb, geom.MoveXY(0, 6), contour.Circle(0.75), n, 2.54)

	for _, rotate := range []geom.Transform{geom.Identity(), geom.RotateD(180)} {
		pcb.With(func() {
			pcb.Transform = pcb.Transform.Multiply(rotate)

			pcb.Track(pads[0], XY{-7.5, -3.2}, XY{-4.5, -3.2}, pins[0])
			pcb.Track(pads[1], XY{-6, -4.2}, XY{-4.8, -4.2}, pins[1])
			pcb.Track(pads[2], XY{pads[2].X, -5.2}, pins[2])
			pcb.Track(pads[3], XY{pads[3].X, -3.5}, pins[3])
			pcb.Track(pads[4], XY{1.5, -4.5}, XY{1.5, -2.5}, pins[4])
			pcb.Track(pads[5], XY{pads[5].X, -2.8}, pins[5])
			pcb.Track(pads[6], XY{pads[6].X, -4.5}, pins[6])
			pcb.Track(pads[7], XY{pads[7].X, -6.2}, pins[7])
			pcb.Track(pads[8], XY{pads[8].X, pins[16].Y}, pins[16])
		})
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
		pcb.SilkText(geom.Move(pads[i]).MoveXY(0, 1.5).RotateD(90).Multiply(textScale), textHeight, loPinNames[i])
		pcb.SilkText(geom.Move(pads[i]).MoveXY(0, 8.3).RotateD(90).Multiply(textScale), textHeight, hiPinNames[i])
	}

	pad.Row(pcb, geom.Identity(), contour.Circle(1.3), 2, 20)
	for x := -10.0; x <= 10; x += 20 {
		t := geom.Move(XY{x, 0})
		pcb.Pad(t, contour.Circle(1.3))
		pcb.Hole(t, contour.Circle(0.8))
	}

	pcb.Cut(contour.RoundRect(24, 16, 2))
}
