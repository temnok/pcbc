package qfn16tiny

import (
	"temnok/pcbc/eda"
	"temnok/pcbc/eda/lib"
	"temnok/pcbc/eda/lib/header/greenconn"
	"temnok/pcbc/eda/lib/pkg/qfn"
	"temnok/pcbc/eda/pcbc"
	"temnok/pcbc/font"
	"temnok/pcbc/path"
	"temnok/pcbc/transform"
)

var (
	mount    = pcbc.MountHole.Arrange(transform.Rotate(30).Move(0, 2.5))
	mountPad = mount.Flatten().Pads.Centers()

	chip = qfn.QFN16G.Arrange(transform.Rotate(-45).Move(0, -1.5))
	pin  = chip.Flatten().Pads.Centers()

	leftLabels  = []string{"PB1", "PA12", "SWD", "SWC", "PF2", "PA0", "PA1", "PA2"}
	rightLabels = []string{"PA8", "VCC", "PB0", "PA7", "PA6", "PA5", "PA4", "PA3"}

	header = &lib.Component{
		Components: lib.Components{
			greenconn.CSCC118(8, false, leftLabels).Arrange(transform.Move(-5.5, 0)),
			greenconn.CSCC118(8, false, rightLabels).Arrange(transform.Move(5.5, 0)),
		},
	}
	pad = header.Flatten().Pads.Centers()

	Board = &lib.Component{
		Cuts: path.Paths{
			path.RoundRect(16, 9.3, 1),
		},

		Components: lib.Components{
			mount,
			chip,
			header,
		},

		Marks: path.Strokes{}.Append(
			font.CenterBold("PY32").
				Apply(transform.Scale(1.7, 1.7).Rotate(-90).Move(2.1, 2)),
			font.CenterBold("F002A").
				Apply(transform.Scale(1.4, 1.7).Rotate(-90).Move(-2.1, 2)),
			pcbc.Logo.Apply(transform.Rotate(-90).Move(2, -3.4)),
		),

		Tracks: path.Strokes{
			0: eda.TrackPaths(
				eda.Track{pin[0]}.YX(pad[0]),
				eda.Track{pin[1]}.YX(pad[1]),
				eda.Track{pin[2]}.YX(pad[2]),
				eda.Track{pin[3]}.YX(pad[3]),
				eda.Track{pin[4]}.YX(pad[4]),
				eda.Track{pin[5]}.DY(-0.3).DX(-2.3).YX(pad[5]),
				eda.Track{pin[6]}.DY(-0.5).DX(-2.5).YX(pad[6]),
				eda.Track{pin[7]}.DY(-0.7).DX(-2.7).YX(pad[7]),

				eda.Track{pin[8]}.DY(-0.7).DX(2.7).YX(pad[15]),
				eda.Track{pin[9]}.DY(-0.5).DX(2.5).YX(pad[14]),
				eda.Track{pin[10]}.DY(-0.3).DX(2.3).YX(pad[13]),
				eda.Track{pin[11]}.YX(pad[12]),
				eda.Track{pin[12]}.YX(pad[11]),
				eda.Track{pin[13]}.YX(pad[10]),
				eda.Track{pin[14]}.YX(pad[9]),
				eda.Track{pin[15]}.DX(2).DY(2.5).YX(pad[8]),
			),
		},

		GroundTracks: path.Strokes{
			0: eda.TrackPaths(
				eda.Track{mountPad[1]}.DY(0.7),
			),
			0.16: eda.TrackPaths(
				eda.Track{pin[16]}.XY(mountPad[4]),
				eda.Track{pin[16]}.DY(-2),
			),
		},
	}
)
