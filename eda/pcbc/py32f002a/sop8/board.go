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

var (
	chip = sop.SOP8.Arrange(geom.MoveXY(2.3, 0))

	pin = chip.Squash().Pads.Centers()

	header = mph100imp40f.G_V_SP_x4.Arrange(geom.MoveXY(0, -6))

	pad = header.Squash().Pads.Centers()

	headerWithTracks = &lib.Component{
		Components: lib.Components{header},
		Tracks: path.Strokes{
			0: eda.TrackPaths(
				eda.Track{pad[1]}.Y(-4.5).X(pin[1].X).Y(-4).Y(pin[1].Y),
				eda.Track{pad[2]}.X(pin[2].X).Y(pin[2].Y),
				eda.Track{pad[3]}.Y(-4).X(pin[3].X).Y(pin[3].Y),
			),
		},
	}

	Board = &lib.Component{
		Cuts: path.Paths{
			path.RoundRect(10.75, 15, 1),
		},

		Components: lib.Components{
			chip,
			headerWithTracks,
			headerWithTracks.Arrange(geom.ScaleXY(1, -1)),
			pcbc.MountHole.Arrange(geom.MoveXY(-2.2, 0)),
		},

		Marks: path.Strokes{}.Append(
			font.CenterBolds([]string{"GND", "PA1", "PA2", "SWD"}, geom.XY{2.54 / 0.9, 0}).
				Transform(geom.MoveXY(0, 4).ScaleXY(0.9, 1.2)),

			font.CenterBold("PY32").Transform(geom.MoveXY(-2.8, 2.4).ScaleXY(1.4, 2)),
			pcbc.Logo.Transform(geom.MoveXY(-4.5, 0).ScaleK(0.8)),
			font.CenterBold("F002A").Transform(geom.MoveXY(-2.8, -2.4).ScaleXY(1.2, 2)),

			font.CenterBolds([]string{"VCC", "PA4", "PA3", "SWC"}, geom.XY{2.54 / 0.9, 0}).
				Transform(geom.MoveXY(0, -4).ScaleXY(0.9, 1.2)),
		),

		Tracks: path.Strokes{
			0: eda.TrackPaths(
				eda.Track{pad[0]}.Y(-4).X(pin[0].X).Y(-3.5).Y(pin[0].Y),
			),
		},

		GroundTracks: path.Strokes{
			0: eda.TrackPaths(
				eda.Track{{pad[0].X, -pad[0].Y}}.Y(4).X(pin[7].X).Y(3.5).Y(pin[7].Y),
				eda.Track{pin[7]}.Y(0).X(-0.4).X(-2.2),
			),
		},
	}
)
