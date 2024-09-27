package pcbc

import (
	"temnok/lab/eda"
	"temnok/lab/eda/lib/pkg/pad"
	"temnok/lab/eda/lib/pkg/qfn16"
	"temnok/lab/font"
	"temnok/lab/geom"
	"temnok/lab/path"
)

type XY = geom.XY

func PY32F002A_QFN16(pcb *eda.PCB, t geom.Transform) {
	pcb.Cut(path.RoundRect(23.5, 11.5, 1).Transform(t))

	textScale := geom.Scale(XY{0.75, 1})
	const titleHeight = 1.5
	pcb.SilkText(t.MoveXY(-5.5, 0).Multiply(textScale), titleHeight, font.Bold, "PY32")
	pcb.SilkText(t.MoveXY(2.3, 0).Multiply(textScale), titleHeight, font.Bold, "F002A")

	qfnT := geom.RotateD(45)
	pins := qfn16.Add(pcb, t.Multiply(qfnT)).Transform(qfnT)

	const tenth = 2.54

	in := path.CutRect(tenth, tenth, 0.3)

	for _, t := range []geom.Transform{t, t.RotateD(180)} {
		padT := geom.MoveXY(0, -4.25)
		pads := pad.Row(pcb, t.Multiply(padT), path.Circle(0.7), in, 9, 2.54, 0).Transform(padT)

		for i := 0; i < 8; i++ {
			pcb.Track(path.Path{pins[i], pads[i]}.Transform(t))
		}
		pcb.Track(path.Path{{0, 0}, {7.5, 0}, pads[8]}.Transform(t))
	}

	pinNames := []string{
		"PB1",
		"PA12",
		"SWD",
		"SWC",
		"PF2",
		"PA0",
		"PA1",
		"PA2",
		"GND",

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

	const textHeight = 1
	textScale = geom.Scale(XY{0.8, 1.2})
	for i := 0; i < 9; i++ {
		pcb.SilkText(t.MoveXY(tenth*float64(i)-11, -2.4).Multiply(textScale), textHeight, font.Bold, pinNames[i])
		pcb.SilkText(t.MoveXY(tenth*float64(i)-11, 2.4).Multiply(textScale), textHeight, font.Bold, pinNames[i+9])
	}

	for x := -7.50; x <= 7.5; x += 15 {
		t := t.Move(XY{x, 0})
		pcb.PadNoStencil(path.Circle(1.3).Transform(t))
		pcb.HoleNoStencil(path.Circle(0.9).Transform(t))

		pcb.StencilHole(path.Pie(8, 1.05, 1.3, 10*geom.Degree).Transform(t)...)
	}
}
