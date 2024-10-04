package py32f002a

import (
	"temnok/lab/eda"
	"temnok/lab/eda/lib"
	"temnok/lab/eda/lib/header/mph100imp40f"
	"temnok/lab/eda/lib/pkg/qfn"
	"temnok/lab/font"
	"temnok/lab/geom"
	"temnok/lab/path"
	"temnok/lab/pcbc"
)

var QFN16 *lib.Component

func init() {
	chipTransform := geom.RotateD(45)
	pins := qfn.QFN16G.Pads.Transform(chipTransform).Centers()

	headerTransform := geom.MoveXY(0, -4.25)
	pads := mph100imp40f.G_V_SP_x9.Pads.Transform(headerTransform).Centers()

	headerWithTracks := &lib.Component{
		Components: lib.Components{
			{
				Transform: headerTransform,
				Components: lib.Components{
					mph100imp40f.G_V_SP_x9,
				},
			},
		},
		Tracks: path.Strokes{
			0: eda.TrackPaths(
				eda.Track{pads[0]}.Y(-2).X(pins[0].X).Y(pins[0].Y),
				eda.Track{pads[1]}.Y(-2.5).X(pins[1].X).Y(pins[1].Y),
				eda.Track{pads[2]}.Y(-3).X(pins[2].X).Y(pins[2].Y),
				eda.Track{pads[3]}.X(-1.25).Y(pins[3].Y).X(pins[3].X),
				eda.Track{pads[4]}.X(1.25).Y(pins[4].Y).X(pins[4].X),
				eda.Track{pads[5]}.Y(pins[5].Y).X(pins[5].X),
				eda.Track{pads[6]}.Y(-2.5).X(pins[6].X).Y(pins[6].Y),
				eda.Track{pads[7]}.Y(-2).X(pins[7].X).Y(pins[7].Y),
				eda.Track{pads[8]}.Y(-2).X(4.5).Y(0).X(pins[16].X),
				eda.Track{{7.5, 0}}.Y(-2),
			),
		},
	}

	board := &lib.Component{
		Description: "PCBC PY32F002A QFN-16 board",

		Components: lib.Components{
			{
				Description: "Chip",
				Transform:   chipTransform,
				Components: lib.Components{
					qfn.QFN16G,
				},
			},
			{
				Description: "Bottom Header",
				Components: lib.Components{
					headerWithTracks,
				},
			},
			{
				Description: "Top Header",
				Transform:   geom.RotateD(180),
				Components: lib.Components{
					headerWithTracks,
				},
			},
			{
				Description: "Left mount hole",
				Transform:   geom.MoveXY(-7.5, 0),
				Components: lib.Components{
					pcbc.MountHole,
				},
			},
			{
				Description: "Right mount hole",
				Transform:   geom.MoveXY(7.5, 0),
				Components: lib.Components{
					pcbc.MountHole,
				},
			},
		},

		Cuts: path.Paths{
			path.RoundRect(23.5, 11.5, 1),
		},

		Marks: path.Strokes{}.Merge(
			font.CenterBold("pc").Transform(geom.MoveXY(-10.6, 0.3).RotateD(45).ScaleK(1.25)),
			font.CenterBold("bc").Transform(geom.MoveXY(-10, -0.3).RotateD(45).ScaleK(1.25)),
			font.CenterBold("TMNK").Transform(geom.MoveXY(10.2, 0.3).ScaleXY(0.75, 0.5)),
			font.CenterBold("TECH").Transform(geom.MoveXY(10.2, -0.3).ScaleXY(0.75, 0.5)),
			font.CenterBold("PY32").Transform(geom.MoveXY(-4.2, 0).ScaleXY(1.3, 2.5)),
			font.CenterBold("F002A").Transform(geom.MoveXY(4.2, 0).ScaleXY(1, 2.5)),
		),

		Tracks: path.Strokes{},
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

	const tenth = 2.54

	for i := 0; i < 9; i++ {
		board.Marks = board.Marks.Merge(
			font.CenterBold(pinNames[i]).Transform(geom.MoveXY(tenth*float64(i-4), -2.4).ScaleXY(0.9, 1.2)),
			font.CenterBold(pinNames[i+9]).Transform(geom.MoveXY(tenth*float64(i-4), 2.4).ScaleXY(0.9, 1.2)),
		)
	}

	QFN16 = board
}
