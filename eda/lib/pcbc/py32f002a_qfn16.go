package pcbc

import (
	"temnok/lab/eda"
	"temnok/lab/eda/lib/pkg/pad"
	"temnok/lab/eda/lib/pkg/qfn16"
	"temnok/lab/geom"
	"temnok/lab/path"
)

type XY = geom.XY

func PY32F002A_QFN16(pcb *eda.PCB, t geom.Transform) {
	pcb.Cut(path.CutRect(18, 10.5, 0.3).Transform(t))

	textScale := geom.Scale(XY{0.75, 1})
	const titleHeight = 1.5
	pcb.SilkText(t.MoveXY(-2.8, -1.5).RotateD(90).Multiply(textScale), titleHeight, "PY32")
	pcb.SilkText(t.MoveXY(2.8, -1.8).RotateD(90).Multiply(textScale), titleHeight, "F002A")

	qfnT := geom.RotateD(45)
	pins := qfn16.Add(pcb, t.Multiply(qfnT)).Transform(qfnT)

	const tenth = 2.54

	//in := path.CutRect(tenth, tenth, 0.3)

	for _, t := range []geom.Transform{t, t.RotateD(180)} {
		ht := geom.MoveXY(0, -1.5*tenth)
		h := pad.Row(pcb, t.Multiply(ht), path.Circle(0.75), nil, 7, 2.54, 0).Transform(ht)

		vt := geom.RotateD(90).MoveXY(0, -3*tenth)
		v := pad.Row(pcb, t.Multiply(vt), path.Circle(0.75), nil, 2, 2.54, 0).Transform(vt)

		pcb.Track(path.Path{pins[0], h[0]}.Transform(t))
		pcb.Track(path.Path{pins[1], h[1]}.Transform(t))
		pcb.Track(path.Path{pins[2], h[2]}.Transform(t))
		pcb.Track(path.Path{pins[3], h[3]}.Transform(t))
		pcb.Track(path.Path{pins[4], h[4]}.Transform(t))
		pcb.Track(path.Path{pins[5], h[5]}.Transform(t))
		pcb.Track(path.Path{pins[6], h[6]}.Transform(t))
		pcb.Track(path.Path{pins[7], {5, -2}, v[0]}.Transform(t))
		pcb.Track(path.Path{{0, 0}, {5, 0}, v[1]}.Transform(t))
	}

	const textHeight = 0.5
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

	textScale = geom.Scale(XY{1.4, 1})
	for i := 0; i < 7; i++ {
		pcb.SilkText(t.MoveXY(1.85*float64(i-2)-2.5, -2.2).Multiply(textScale), textHeight, loPinNames[i])
		pcb.SilkText(t.MoveXY(4.9, -1.6).Multiply(textScale), textHeight, loPinNames[7])
		pcb.SilkText(t.MoveXY(4.9, 1.6).Multiply(textScale), textHeight, loPinNames[8])

		pcb.SilkText(t.MoveXY(1.85*float64(i-2)-2.5, 2.2).Multiply(textScale), textHeight, hiPinNames[i+2])
		pcb.SilkText(t.MoveXY(-6.2, 1.6).Multiply(textScale), textHeight, hiPinNames[1])
		pcb.SilkText(t.MoveXY(-6.2, -1.6).Multiply(textScale), textHeight, hiPinNames[0])
	}

	for x := -5.0; x <= 5; x += 10 {
		t := t.Move(XY{x, 0})
		pcb.PadNoStencil(path.Circle(1.3).Transform(t))
		pcb.HoleNoStencil(path.Circle(0.9).Transform(t))

		pcb.StencilHole(path.Pie(8, 1.05, 1.3, 10*geom.Degree).Transform(t)...)
	}
}
