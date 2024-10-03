package pcbc

import (
	"temnok/lab/eda"
	"temnok/lab/eda/lib"
	"temnok/lab/eda/lib/header/mph100imp40f"
	"temnok/lab/eda/lib/pkg/qfn"
	"temnok/lab/font"
	"temnok/lab/geom"
	"temnok/lab/path"
)

func PY32F002A_QFN16() *lib.Component {
	comp := &lib.Component{
		Cuts: path.Paths{path.RoundRect(23.5, 11.5, 1)},
		Marks: path.Strokes{}.Merge(
			font.CenterBold("pc").Transform(geom.MoveXY(-10.6, 0.3).RotateD(45).ScaleK(1.25)),
			font.CenterBold("bc").Transform(geom.MoveXY(-10, -0.3).RotateD(45).ScaleK(1.25)),
			font.CenterBold("TMNK").Transform(geom.MoveXY(10.2, 0.3).ScaleXY(0.75, 0.5)),
			font.CenterBold("TECH").Transform(geom.MoveXY(10.2, -0.3).ScaleXY(0.75, 0.5)),
			font.CenterBold("PY32").Transform(geom.MoveXY(-4, 0).ScaleXY(1.3, 2.5)),
			font.CenterBold("F002A").Transform(geom.MoveXY(4, 0).ScaleXY(1, 2.5)),
		),
	}

	chip := qfn.QFN16G().Transform(geom.RotateD(45))
	comp = comp.Merge(chip)
	pins := chip.Pads.Centers()

	const tenth = 2.54

	for _, dt := range []geom.Transform{geom.Identity(), geom.RotateD(180)} {
		header := mph100imp40f.Gvsp(9).Transform(geom.MoveXY(0, -4.25))
		comp = comp.Merge(header.Transform(dt))

		pads := header.Pads.Centers()

		comp.Tracks[0] = append(comp.Tracks[0], eda.TrackPaths(
			eda.Track{pads[0]}.Y(-2).X(pins[0].X).Y(pins[0].Y),
			eda.Track{pads[1]}.Y(-2.5).X(pins[1].X).Y(pins[1].Y),
			eda.Track{pads[2]}.Y(-3).X(pins[2].X).Y(pins[2].Y),
			eda.Track{pads[3]}.X(-1.25).Y(pins[3].Y).X(pins[3].X),
			eda.Track{pads[4]}.X(1.25).Y(pins[4].Y).X(pins[4].X),
			eda.Track{pads[5]}.Y(pins[5].Y).X(pins[5].X),
			eda.Track{pads[6]}.Y(-2.5).X(pins[6].X).Y(pins[6].Y),
			eda.Track{pads[7]}.Y(-2).X(pins[7].X).Y(pins[7].Y),
			eda.Track{{0, 0}}.X(7.5).X(pads[8].X).Y(pads[8].Y),
		).Transform(dt)...)
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
		comp.Marks = comp.Marks.Merge(
			font.CenterBold(pinNames[i]).Transform(geom.MoveXY(tenth*float64(i-4), -2.4).ScaleXY(0.9, 1.2)),
			font.CenterBold(pinNames[i+9]).Transform(geom.MoveXY(tenth*float64(i-4), 2.4).ScaleXY(0.9, 1.2)),
		)
	}

	for x := -7.50; x <= 7.5; x += 15 {
		dt := geom.MoveXY(x, 0)
		outC := path.Circle(2.6).Transform(dt)
		inC := path.Circle(1.8).Transform(dt)
		comp.Openings = append(comp.Openings, outC)
		comp.Pads = append(comp.Pads, outC)
		comp.Pads = append(comp.Pads, path.Pie(6, 1.0, 1.3, 10*geom.Degree).Transform(dt)...)
		comp.Pads = append(comp.Pads, path.Pie(6, 1.0, 1.3, 10*geom.Degree).Transform(dt)...)
		comp.Holes = append(comp.Holes, inC)
	}

	return comp
}
