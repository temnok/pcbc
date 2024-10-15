package sop8

import (
	"temnok/pcbc/eda"
	"temnok/pcbc/eda/lib"
	"temnok/pcbc/eda/lib/header/mph100imp40f"
	"temnok/pcbc/eda/lib/pkg/sop"
	"temnok/pcbc/eda/pcbc"
	"temnok/pcbc/font"
	"temnok/pcbc/geom"
	"temnok/pcbc/path"
)

var Board *lib.Component

func init() {
	chipTransform := geom.MoveXY(2.3, 0)
	pins := sop.SOP8.Pads.Transform(chipTransform).Centers()

	headerTransform := geom.MoveXY(0, -6)
	pads := mph100imp40f.G_V_SP_x4.Pads.Transform(headerTransform).Centers()

	headerWithTracks := &lib.Component{
		Components: lib.Components{
			{
				Transform: headerTransform,
				Components: lib.Components{
					mph100imp40f.G_V_SP_x4,
				},
			},
		},
		Tracks: path.Strokes{
			0: eda.TrackPaths(
				eda.Track{pads[1]}.Y(-4.5).X(pins[1].X).Y(-4).Y(pins[1].Y),
				eda.Track{pads[2]}.X(pins[2].X).Y(pins[2].Y),
				eda.Track{pads[3]}.Y(-4).X(pins[3].X).Y(pins[3].Y),
			),
		},
	}

	board := &lib.Component{
		Description: "PCBC PY32F002A SOP-8 board",

		Components: lib.Components{
			{
				Description: "Chip",
				Transform:   chipTransform,
				Components: lib.Components{
					sop.SOP8,
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
				Transform:   geom.ScaleXY(1, -1),
				Components: lib.Components{
					headerWithTracks,
				},
			},
			{
				Description: "Mount hole",
				Transform:   geom.MoveXY(-2.2, 0),
				Components: lib.Components{
					pcbc.MountHole,
				},
			},
		},

		Cuts: path.Paths{
			path.RoundRect(10.75, 15, 1),
		},

		Marks: path.Strokes{}.Append(
			font.CenterBold("PY32").Transform(geom.MoveXY(-2.8, 2.4).ScaleXY(1.4, 2)),
			font.CenterBold("F002A").Transform(geom.MoveXY(-2.8, -2.4).ScaleXY(1.2, 2)),
			pcbc.Logo.Transform(geom.MoveXY(-4.5, 0).ScaleK(0.8)),
		),

		Tracks: path.Strokes{
			0: eda.TrackPaths(
				eda.Track{pads[0]}.Y(-4).X(pins[0].X).Y(-3.5).Y(pins[0].Y),
			),
		},

		GroundTracks: path.Strokes{
			0: eda.TrackPaths(
				eda.Track{{pads[0].X, -pads[0].Y}}.Y(4).X(pins[7].X).Y(3.5).Y(pins[7].Y),
				eda.Track{pins[7]}.Y(0).X(-0.4).X(-2.2),
			),
		},
	}

	pinNames := []string{
		"VCC",
		"PA4",
		"PA3",
		"SWC",
		"SWD",
		"PA2",
		"PA1",
		"GND",
	}

	const tenth = 2.54

	for i := 0; i < 4; i++ {
		board.Marks = board.Marks.Append(
			font.CenterBold(pinNames[i]).Transform(geom.MoveXY(tenth*(float64(i)-1.5), -4).ScaleXY(0.9, 1.2)),
			font.CenterBold(pinNames[7-i]).Transform(geom.MoveXY(tenth*(float64(i)-1.5), 4).ScaleXY(0.9, 1.2)),
		)
	}

	Board = board
}
