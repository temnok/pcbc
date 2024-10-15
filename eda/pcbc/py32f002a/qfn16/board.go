package qfn16

import (
	"temnok/pcbc/eda"
	"temnok/pcbc/eda/lib"
	"temnok/pcbc/eda/lib/header/mph100imp40f"
	"temnok/pcbc/eda/lib/pkg/qfn"
	"temnok/pcbc/eda/pcbc"
	"temnok/pcbc/font"
	"temnok/pcbc/geom"
	"temnok/pcbc/path"
)

var (
	chipTransform = geom.RotateD(45)
	pins          = qfn.QFN16G.Pads.Transform(chipTransform).Centers()

	headerTransform = geom.MoveXY(0, -4.25)
	pads            = mph100imp40f.G_V_SP_x8.Pads.Transform(headerTransform).Centers()

	headerWithTracks = &lib.Component{
		Components: lib.Components{
			{
				Transform: headerTransform,
				Components: lib.Components{
					mph100imp40f.G_V_SP_x8,
				},
			},
		},
		Tracks: path.Strokes{
			0: eda.TrackPaths(
				eda.Track{pads[0]}.Y(-2.5).X(-4.9).Y(-2).XY(pins[0]),
				eda.Track{pads[1]}.Y(-2.5).XY(pins[1]),
				eda.Track{pads[2]}.YX(pins[2]),
				eda.Track{pads[3]}.X(-1.25).YX(pins[3]),
				eda.Track{pads[4]}.X(1.25).YX(pins[4]),
				eda.Track{pads[5]}.YX(pins[5]),
				eda.Track{pads[6]}.Y(-2.5).XY(pins[6]),
				eda.Track{pads[7]}.Y(-2.5).X(4.9).Y(-2).XY(pins[7]),
			),
		},
		GroundTracks: path.Strokes{
			0.16: eda.TrackPaths(
				eda.Track{{9.5, 0}}.X(pins[16].X),
			),
		},
	}

	Board = &lib.Component{
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
				Description: "Top Header",
				Transform:   geom.RotateD(180),
				Components: lib.Components{
					headerWithTracks,
				},
			},
			{
				Description: "Bottom Header",
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
			path.RoundRect(21, 11.5, 1),
		},

		Marks: path.Strokes{}.Append(
			font.CenterBolds([]string{
				"PA8",
				"VCC",
				"PB0",
				"PA7",
				"PA6",
				"PA5",
				"PA4",
				"PA3",
			}, geom.XY{2.54 / 0.9, 0}).Transform(geom.MoveXY(0, 2.4).ScaleXY(0.9, 1.2)),

			font.CenterBolds([]string{
				"PB1",
				"PA12",
				"SWD",
				"SWC",
				"PF2",
				"PA0",
				"PA1",
				"PA2",
			}, geom.XY{2.54 / 0.9, 0}).Transform(geom.MoveXY(0, -2.4).ScaleXY(0.9, 1.2)),

			pcbc.Logo.Transform(geom.MoveXY(-9.7, 0).ScaleK(0.8)),
			pcbc.TmnkTech.Transform(geom.MoveXY(9.7, 0).RotateD(90)),
			font.CenterBold("PY32").Transform(geom.MoveXY(-4.2, 0).ScaleXY(1.3, 2.5)),
			font.CenterBold("F002A").Transform(geom.MoveXY(4.2, 0).ScaleXY(1, 2.5)),
		),

		Tracks: path.Strokes{},
	}
)
