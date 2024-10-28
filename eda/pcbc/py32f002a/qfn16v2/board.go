package qfn16v2

import (
	"temnok/pcbc/eda"
	"temnok/pcbc/eda/lib/header/greenconn"
	"temnok/pcbc/eda/lib/pkg/qfn"
	"temnok/pcbc/eda/pcbc"
	"temnok/pcbc/font"
	"temnok/pcbc/path"
	"temnok/pcbc/transform"
)

var (
	chip     = qfn.QFN16G.Arrange(transform.Rotate(-45).Move(0, 2))
	pin      = chip.PadCenters()
	pinNames = []string{"PB1", "PA12", "SWD", "SWC", "PF2", "PA0", "PA1", "PA2", "GND",
		"PA3", "PA4", "PA5", "PA6", "PA7", "PB0", "VCC", "PA8"}

	header = &eda.Component{
		Components: eda.Components{
			greenconn.CSCC118(17, false, pinNames).Arrange(transform.Rotate(90).Move(0, -2.5)),
		},
	}
	pad = header.PadCenters()

	Board = &eda.Component{
		Cuts: path.Paths{
			path.RoundRect(19, 10, 1),
		},

		Components: eda.Components{
			pcbc.MountHole.Arrange(transform.Rotate(45).Move(-7.5, 3)),
			pcbc.MountHole.Arrange(transform.Rotate(-45).Move(7.5, 3)),
			chip,
			header,
		},

		Tracks: eda.TrackPaths(
			eda.Track{pin[0]}.DY(0.7).DX(-2.4).XY(pad[0]),
			eda.Track{pin[1]}.DY(0.5).DX(-1.7).XY(pad[1]),
			eda.Track{pin[2]}.DY(0.3).DX(-1).XY(pad[2]),
			eda.Track{pin[3]}.DY(0.1).DX(-0.3).XY(pad[3]),
			eda.Track{pin[4]}.XY(pad[4]),
			eda.Track{pin[5]}.XY(pad[5]),
			eda.Track{pin[6]}.XY(pad[6]),
			eda.Track{pin[7]}.XY(pad[7]),

			eda.Track{pin[8]}.XY(pad[9]),
			eda.Track{pin[9]}.XY(pad[10]),
			eda.Track{pin[10]}.XY(pad[11]),
			eda.Track{pin[11]}.XY(pad[12]),
			eda.Track{pin[12]}.DY(0.1).DX(0.3).XY(pad[13]),
			eda.Track{pin[13]}.DY(0.3).DX(1).XY(pad[14]),
			eda.Track{pin[14]}.DY(0.5).DX(1.7).XY(pad[15]),
			eda.Track{pin[15]}.DY(0.7).DX(2.4).XY(pad[16]),
		),

		GroundTracks: eda.TrackPaths(
			eda.Track{pad[8]}.DY(1.5).DY(-3),
		),

		Marks: path.Strokes{}.Append(
			font.CenterBold("PY").Apply(transform.ScaleK(2.5).Move(-4, 3.4)),
			font.CenterBold("32").Apply(transform.ScaleK(2.5).Move(-4, 1)),
			font.CenterBold("F00").Apply(transform.Scale(2, 2.5).Move(4, 3.4)),
			font.CenterBold("2A").Apply(transform.ScaleK(2.5).Move(4, 1)),
			pcbc.Logo.Apply(transform.Move(-8.3, 0.8)),
		),
	}
)
