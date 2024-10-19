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
	"temnok/pcbc/transform"
)

var (
	chip = sop.SOP8.Arrange(transform.Move(2.3, 0))

	pin = chip.Squash().Pads.Centers()

	header = mph100imp40f.G_V_SP(4).Arrange(transform.Move(0, -6))

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
			headerWithTracks.Arrange(transform.Scale(1, -1)),
			pcbc.MountHole.Arrange(transform.Move(-2.2, 0)),
		},

		Marks: path.Strokes{}.Append(
			font.CenterBolds([]string{"GND", "PA1", "PA2", "SWD"}, geom.XY{2.54 / 0.9, 0}).
				Apply(transform.Scale(0.9, 1.2).Move(0, 4)),

			font.CenterBold("PY32").Apply(transform.Scale(1.4, 2).Move(-2.8, 2.4)),
			pcbc.Logo.Apply(transform.ScaleK(0.8).Move(-4.5, 0)),
			font.CenterBold("F002A").Apply(transform.Scale(1.2, 2).Move(-2.8, -2.4)),

			font.CenterBolds([]string{"VCC", "PA4", "PA3", "SWC"}, geom.XY{2.54 / 0.9, 0}).
				Apply(transform.Scale(0.9, 1.2).Move(0, -4)),
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
