package qfn16

import (
	"temnok/pcbc/eda"
	"temnok/pcbc/eda/lib/header/mph100imp40f"
	"temnok/pcbc/eda/lib/pkg/qfn"
	"temnok/pcbc/eda/pcbc"
	"temnok/pcbc/font"
	"temnok/pcbc/path"
	"temnok/pcbc/transform"
)

var (
	chip = qfn.QFN16G.Arrange(transform.Rotate(45))

	pin = chip.PadCenters()

	header = mph100imp40f.G_V_SP(8).Arrange(transform.Move(0, -4.25))

	pad = header.PadCenters()

	headerWithTracks = &eda.Component{
		Components: eda.Components{header},
		Tracks: eda.TrackPaths(
			eda.Track{pad[0]}.Y(-2.5).X(-4.9).Y(-2).XY(pin[0]),
			eda.Track{pad[1]}.Y(-2.5).XY(pin[1]),
			eda.Track{pad[2]}.YX(pin[2]),
			eda.Track{pad[3]}.X(-1.25).YX(pin[3]),
			eda.Track{pad[4]}.X(1.25).YX(pin[4]),
			eda.Track{pad[5]}.YX(pin[5]),
			eda.Track{pad[6]}.Y(-2.5).XY(pin[6]),
			eda.Track{pad[7]}.Y(-2.5).X(4.9).Y(-2).XY(pin[7]),
		),
	}

	labelScale = transform.Scale(0.9, 1.2)

	Board = &eda.Component{
		Components: eda.Components{
			chip,
			headerWithTracks.Arrange(transform.Rotate(180)),
			headerWithTracks,
			pcbc.MountHole.Arrange(transform.Move(-7.5, 0)),
			pcbc.MountHole.Arrange(transform.Move(7.5, 0)),
		},

		Cuts: path.Paths{
			path.RoundRect(21, 11.5, 1),
		},

		Marks: path.Strokes{}.Append(
			font.CenterBolds([]string{"PA8", "VCC", "PB0", "PA7", "PA6", "PA5", "PA4", "PA3"},
				path.Point{X: 2.54 / 0.9}).Apply(labelScale.Move(0, 2.4)),

			pcbc.Logo.Apply(transform.ScaleK(0.8).Move(-9.7, 0)),
			font.CenterBold("PY32").Apply(transform.Scale(1.3, 2.5).Move(-4.2, 0)),
			font.CenterBold("F002A").Apply(transform.Scale(1, 2.5).Move(4.2, 0)),
			pcbc.TmnkTech.Apply(transform.Rotate(90).Move(9.7, 0)),

			font.CenterBolds([]string{"PB1", "PA12", "SWD", "SWC", "PF2", "PA0", "PA1", "PA2"},
				path.Point{X: 2.54 / 0.9}).Apply(labelScale.Move(0, -2.4)),
		),
	}
)
