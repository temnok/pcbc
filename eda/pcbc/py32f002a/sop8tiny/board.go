package sop8tiny

import (
	"temnok/pcbc/eda"
	"temnok/pcbc/eda/lib"
	"temnok/pcbc/eda/lib/header/greenconn"
	"temnok/pcbc/eda/lib/pkg/sop"
	"temnok/pcbc/eda/pcbc"
	"temnok/pcbc/font"
	"temnok/pcbc/path"
	"temnok/pcbc/transform"
)

var (
	mount    = pcbc.MountHole.Arrange(transform.Move(-4.8, 0))
	mountPad = mount.Squash().Pads.Centers()

	chip = sop.SOP8.Arrange(transform.Move(-0.3, 0))
	pin  = chip.Squash().Pads.Centers()

	header = greenconn.CSCC118(7, []string{"SWD", "PA2", "PA1", "VCC", "PA4", "PA3", "SWC"}).
		Arrange(transform.Move(5.2, 0))
	pad = header.Squash().Pads.Centers()

	Board = &lib.Component{
		Cuts: path.Paths{
			path.RoundRect(16, 8, 1),
		},

		Components: lib.Components{
			pcbc.MountHole.Arrange(transform.Move(-4.8, 0)),
			chip,
			header,
		},

		Marks: path.Strokes{}.Append(
			font.CenterBold("PY32").Apply(transform.Scale(1.5, 2.8).Move(-2.8, 2.7)),
			pcbc.Logo.Apply(transform.Move(-4.5, 0)),
			//pcbc.TmnkTech.Apply(transform.ScaleK(0.8).Move(-4.4, -0.9)),
			font.CenterBold("F002A").Apply(transform.Scale(1.3, 2.8).Move(-2.8, -2.7)),
		).Apply(transform.Move(-2.6, 0)),

		Tracks: path.Strokes{
			0: eda.TrackPaths(
				eda.Track{pin[0]}.DY(0.8).YX(pad[3]),
				eda.Track{pin[1]}.DY(0.8).YX(pad[4]),
				eda.Track{pin[2]}.DY(0.8).DY(0.3).DX(2.5).YX(pad[5]),
				eda.Track{pin[3]}.YX(pad[6]),
				eda.Track{pin[4]}.YX(pad[0]),
				eda.Track{pin[5]}.DY(-0.8).DY(-0.3).DX(2.5).YX(pad[1]),
				eda.Track{pin[6]}.DY(-0.8).YX(pad[2]),
			),
		},

		GroundTracks: path.Strokes{
			0: eda.TrackPaths(
				eda.Track{pin[7]}.DY(0.7),
				eda.Track{pin[7]}.DY(-2).YX(mountPad[0]),
				eda.Track{mountPad[3]}.DX(-0.7),
			),
		},
	}
)
