package sop8

import (
	"temnok/pcbc/eda"
	"temnok/pcbc/eda/lib/header/mph100imp40f"
	"temnok/pcbc/eda/lib/pkg/sop"
	"temnok/pcbc/eda/pcbc"
	"temnok/pcbc/font"
	"temnok/pcbc/path"
	"temnok/pcbc/transform"
)

var (
	chip = sop.SOP8.Arrange(transform.Move(2.3, 0))

	pin = chip.PadCenters()

	header = mph100imp40f.G_V_SP(4).Arrange(transform.Move(0, -6))

	pad = header.PadCenters()

	headerWithTracks = &eda.Component{
		Components: eda.Components{header},
		Tracks: eda.TrackPaths(
			eda.Track{pad[1]}.Y(-4.5).X(pin[1].X).Y(-4).Y(pin[1].Y),
			eda.Track{pad[2]}.X(pin[2].X).Y(pin[2].Y),
			eda.Track{pad[3]}.Y(-4).X(pin[3].X).Y(pin[3].Y),
		),
	}

	Board = &eda.Component{
		Cuts: path.Paths{
			path.RoundRect(10.75, 15, 1),
		},

		Components: eda.Components{
			chip,
			headerWithTracks,
			headerWithTracks.Arrange(transform.Scale(1, -1)),
			pcbc.MountHole.Arrange(transform.Rotate(90).Move(-2.2, 0)),

			{
				Transform: transform.Scale(0.9, 1.2).Move(0, 4),
				Marks: font.ShiftedCenteredPaths(path.Point{X: 2.54 / 0.9},
					"GND", "PA1", "PA2", "SWD"),
			},

			{
				Transform: transform.Scale(0.9, 1.2).Move(0, -4),
				Marks: font.ShiftedCenteredPaths(path.Point{X: 2.54 / 0.9},
					"VCC", "PA4", "PA3", "SWC"),
			},

			eda.CenteredText("PY32").Arrange(transform.Scale(1.4, 2).Move(-2.8, 2.4)),

			eda.CenteredText("F002A").Arrange(transform.Scale(1.2, 2).Move(-2.8, -2.4)),

			pcbc.Logo.Arrange(transform.ScaleK(0.8).Move(-4.5, 0)),
		},

		Tracks: eda.TrackPaths(
			eda.Track{pad[0]}.Y(-4).X(pin[0].X).Y(-3.5).Y(pin[0].Y),
		),

		GroundTracks: eda.TrackPaths(
			eda.Track{pin[7]}.DY(-1).DY(2),
		),
	}
)
