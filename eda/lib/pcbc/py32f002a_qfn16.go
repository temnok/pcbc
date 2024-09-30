package pcbc

import (
	"temnok/lab/eda"
	"temnok/lab/eda/lib/header/mph100imp40f"
	"temnok/lab/eda/lib/pkg/qfn"
	"temnok/lab/geom"
	"temnok/lab/path"
)

type XY = geom.XY

func PY32F002A_QFN16(pcb *eda.PCB, t geom.Transform) {
	pcb.Cut(path.RoundRect(23.5, 11.5, 1).Transform(t))

	pcb.SilkText(t.MoveXY(-10.6, 0.3).RotateD(45).ScaleK(1.25), "pc")
	pcb.SilkText(t.MoveXY(-10, -0.3).RotateD(45).ScaleK(1.25), "bc")
	pcb.SilkText(t.MoveXY(10.2, 0.3).ScaleXY(0.75, 0.5), "TMNK")
	pcb.SilkText(t.MoveXY(10.2, -0.3).ScaleXY(0.75, 0.5), "TECH")

	pcb.SilkText(t.MoveXY(-4, 0).ScaleXY(1.3, 2.5), "PY32")
	pcb.SilkText(t.MoveXY(4, 0).ScaleXY(1, 2.5), "F002A")

	chip := qfn.QFN16G().Transform(geom.RotateD(45))

	pcb.Component(chip.Transform(t))
	pins := chip.Pads.Centers()

	const tenth = 2.54

	//ins := path.CutRect(tenth, tenth, 0.3)

	for _, t := range []geom.Transform{t, t.RotateD(180)} {
		header := mph100imp40f.Gvsp(9).Transform(geom.MoveXY(0, -4.25))
		pcb.Component(header.Transform(t))

		pads := header.Pads.Centers()

		pcb.Track(eda.Track{pads[0]}.Y(-2).X(pins[0].X).Y(pins[0].Y).Transform(t))
		pcb.Track(eda.Track{pads[1]}.Y(-2.5).X(pins[1].X).Y(pins[1].Y).Transform(t))
		pcb.Track(eda.Track{pads[2]}.Y(-3).X(pins[2].X).Y(pins[2].Y).Transform(t))
		pcb.Track(eda.Track{pads[3]}.X(-1.25).Y(pins[3].Y).X(pins[3].X).Transform(t))
		pcb.Track(eda.Track{pads[4]}.X(1.25).Y(pins[4].Y).X(pins[4].X).Transform(t))
		pcb.Track(eda.Track{pads[5]}.Y(pins[5].Y).X(pins[5].X).Transform(t))
		pcb.Track(eda.Track{pads[6]}.Y(-2.5).X(pins[6].X).Y(pins[6].Y).Transform(t))
		pcb.Track(eda.Track{pads[7]}.Y(-2).X(pins[7].X).Y(pins[7].Y).Transform(t))

		pcb.Track(eda.Track{{0, 0}}.X(7.5).X(pads[8].X).Y(pads[8].Y).Transform(t))
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

	for i := 0; i < 9; i++ {
		pcb.SilkText(t.MoveXY(tenth*float64(i-4), -2.4).ScaleXY(0.9, 1.2), pinNames[i])
		pcb.SilkText(t.MoveXY(tenth*float64(i-4), 2.4).ScaleXY(0.9, 1.2), pinNames[i+9])
	}

	for x := -7.50; x <= 7.5; x += 15 {
		t := t.MoveXY(x, 0)

		outC := path.Circle(1.3).Transform(t)
		pcb.PadNoStencil(outC)
		pcb.HoleNoStencil(path.Circle(0.9).Transform(t))

		pcb.StencilHole(path.Pie(8, 1.05, 1.3, 10*geom.Degree).Transform(t)...)

		pcb.PlacerHole(outC)
	}
}
