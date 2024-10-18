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
	chip = qfn.QFN16G.Arrange(geom.RotateD(45))

	pin = chip.Squash().Pads.Centers()

	header = mph100imp40f.G_V_SP_x8.Arrange(geom.MoveXY(0, -4.25))

	pad = header.Squash().Pads.Centers()

	headerWithTracks = &lib.Component{
		Components: lib.Components{header},
		Tracks: path.Strokes{
			0: eda.TrackPaths(
				eda.Track{pad[0]}.Y(-2.5).X(-4.9).Y(-2).XY(pin[0]),
				eda.Track{pad[1]}.Y(-2.5).XY(pin[1]),
				eda.Track{pad[2]}.YX(pin[2]),
				eda.Track{pad[3]}.X(-1.25).YX(pin[3]),
				eda.Track{pad[4]}.X(1.25).YX(pin[4]),
				eda.Track{pad[5]}.YX(pin[5]),
				eda.Track{pad[6]}.Y(-2.5).XY(pin[6]),
				eda.Track{pad[7]}.Y(-2.5).X(4.9).Y(-2).XY(pin[7]),
			),
		},
		GroundTracks: path.Strokes{
			0.16: eda.TrackPaths(
				eda.Track{{9.5, 0}}.X(pin[16].X),
			),
		},
	}

	Board = &lib.Component{
		Components: lib.Components{
			chip,
			headerWithTracks.Arrange(geom.RotateD(180)),
			headerWithTracks,
			pcbc.MountHole.Arrange(geom.MoveXY(-7.5, 0)),
			pcbc.MountHole.Arrange(geom.MoveXY(7.5, 0)),
		},

		Cuts: path.Paths{
			path.RoundRect(21, 11.5, 1),
		},

		Marks: path.Strokes{}.Append(
			font.CenterBolds([]string{"PA8", "VCC", "PB0", "PA7", "PA6", "PA5", "PA4", "PA3"},
				geom.XY{2.54 / 0.9, 0}).Transform(geom.MoveXY(0, 2.4).ScaleXY(0.9, 1.2)),

			pcbc.Logo.Transform(geom.MoveXY(-9.7, 0).ScaleK(0.8)),
			font.CenterBold("PY32").Transform(geom.MoveXY(-4.2, 0).ScaleXY(1.3, 2.5)),
			font.CenterBold("F002A").Transform(geom.MoveXY(4.2, 0).ScaleXY(1, 2.5)),
			pcbc.TmnkTech.Transform(geom.MoveXY(9.7, 0).RotateD(90)),

			font.CenterBolds([]string{"PB1", "PA12", "SWD", "SWC", "PF2", "PA0", "PA1", "PA2"},
				geom.XY{2.54 / 0.9, 0}).Transform(geom.MoveXY(0, -2.4).ScaleXY(0.9, 1.2)),
		),
	}
)
