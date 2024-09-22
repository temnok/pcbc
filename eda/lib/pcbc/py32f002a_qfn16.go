package pcbc

import (
	"temnok/lab/contour"
	"temnok/lab/eda"
	"temnok/lab/eda/lib/pkg/pad"
	"temnok/lab/eda/lib/pkg/qfn16"
	"temnok/lab/geom"
)

type XY = geom.XY

func PY32F002A_QFN16(pcb *eda.PCB, t geom.Transform) {
	pcb.Cut(t.Points(contour.RoundRect(24, 17, 1.5)))

	textScale := geom.Scale(XY{0.75, 1})
	titleHeight := 2.0
	textHeight := 1.5

	pcb.SilkText(t.MoveXY(-7, 0).Multiply(textScale), titleHeight, "PY32")
	pcb.SilkText(t.MoveXY(2.75, 0).Multiply(textScale), titleHeight, "F002A")

	qfnT := geom.RotateD(45)
	pins := qfnT.Points(qfn16.Add(pcb, t.Multiply(qfnT)))

	n := 9

	padT := geom.MoveXY(0, -6)
	pads := padT.Points(pad.Row(pcb, t.Multiply(padT), contour.Circle(0.75), n, 2.54, 0))
	pad.Row(pcb, t.MoveXY(0, 6), contour.Circle(0.75), n, 2.54, 0)

	//padT := geom.MoveXY(0, -5)
	//padC := contour.Rect(0.8, 3.1)
	//pads := padT.Points(pad.Row(pcb, t.Multiply(padT), padC, n, 2.54, 1.3))
	//pad.Row(pcb, t.MoveXY(0, 5), padC, n, 2.54, -1.3)

	for _, t := range []geom.Transform{t, t.RotateD(180)} {
		pcb.Track(t.Points([]XY{pads[0], {-7.5, -3.2}, {-4.5, -3.2}, pins[0]}))
		pcb.Track(t.Points([]XY{pads[1], {-6, -4.2}, {-4.8, -4.2}, pins[1]}))
		pcb.Track(t.Points([]XY{pads[2], {-3.8, -4.5}, {-3.8, -4}, pins[2]}))
		pcb.Track(t.Points([]XY{pads[3], {pads[3].X, -3.5}, pins[3]}))
		pcb.Track(t.Points([]XY{pads[4], {1.5, -4.5}, {1.5, -2.5}, pins[4]}))
		pcb.Track(t.Points([]XY{pads[5], {pads[5].X, -2.8}, pins[5]}))
		pcb.Track(t.Points([]XY{pads[6], {pads[6].X, -4.5}, pins[6]}))
		pcb.Track(t.Points([]XY{pads[7], {pads[7].X, -6.2}, pins[7]}))
		pcb.Track(t.Points([]XY{pads[8], {pads[8].X, pins[16].Y}, pins[16]}))
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

	for x := -10.0; x <= 10; x += 20 {
		t := t.Move(XY{x, 0})
		pcb.PadNoStencil(t.Points(contour.Circle(1.35)))
		pcb.HoleNoStencil(t.Points(contour.Circle(0.85)))

		pcb.StencilHole(t.PointsAll(contour.Pie(8, 1, 1.25, 10*geom.Degree))...)
	}
}
