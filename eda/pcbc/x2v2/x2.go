package x2v2

import (
	"temnok/pcbc/eda"
	"temnok/pcbc/eda/lib"
	"temnok/pcbc/eda/lib/header/greenconn"
	"temnok/pcbc/eda/lib/pkg/smd"
	"temnok/pcbc/eda/pcbc"
	"temnok/pcbc/path"
	"temnok/pcbc/transform"
)

func X2(topLabel, bottomLabel string) *lib.Component {
	header := greenconn.CSCC118(3, false, []string{topLabel, "GND", bottomLabel}).Arrange(transform.Move(-2, 0))
	pad := header.Flatten().Pads.Centers()

	chip := smd.I0603.Arrange(transform.Rotate(-90).Move(0.8, 0))
	pin := chip.Flatten().Pads.Centers()

	mount := pcbc.MountHole.Arrange(transform.Move(3, 0))
	sink := mount.Flatten().Pads.Centers()

	return &lib.Component{
		Cuts: path.Paths{
			path.RoundRect(10, 4, 1),
		},

		Components: lib.Components{
			header,
			chip,
			mount,
		},

		Tracks: path.Strokes{
			0: eda.TrackPaths(
				eda.Track{pad[0]}.XY(pin[0]),
				eda.Track{pad[2]}.XY(pin[1]),
			),
		},

		GroundTracks: path.Strokes{
			0: eda.TrackPaths(
				eda.Track{pad[1]}.DX(-1.2),
				eda.Track{pad[1]}.XY(sink[3]),
				eda.Track{sink[2]}.DXY(-0.5, 0.5),
				eda.Track{sink[4]}.DXY(-0.5, -0.5),
			),
		},
	}
}
