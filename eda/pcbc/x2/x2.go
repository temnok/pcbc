// Copyright © 2025 Alex Temnok. All rights reserved.

package x2

import (
	"temnok/pcbc/eda"
	"temnok/pcbc/eda/lib/header/greenconn"
	"temnok/pcbc/eda/lib/pkg/smd"
	"temnok/pcbc/eda/pcbc"
	"temnok/pcbc/path"
	"temnok/pcbc/transform"
)

func X2(topLabel, bottomLabel string) *eda.Component {
	header := greenconn.CSCC118(3, false, []string{topLabel, "GND", bottomLabel}).Arrange(transform.Move(-2, 0))
	pad := header.PadCenters()

	chip := smd.I0603.Arrange(transform.Rotate(-90).Move(0.8, 0))
	pin := chip.PadCenters()

	mount := pcbc.MountHole.Arrange(transform.Move(3, 0))
	sink := mount.PadCenters()

	return &eda.Component{
		Cuts: path.RoundRect(10, 4, 1),

		Components: eda.Components{
			header,
			chip,
			mount,
		},

		Tracks: eda.Tracks(
			eda.Track{pad[0]}.XY(pin[0]),
			eda.Track{pad[2]}.XY(pin[1]),
		),

		GroundTracks: eda.Tracks(
			eda.Track{pad[1]}.DX(-1.2),
			eda.Track{pad[1]}.XY(sink[3]),
			eda.Track{sink[2]}.DX(-0.5).DY(0.5),
			eda.Track{sink[4]}.DX(-0.5).DY(-0.5),
		),
	}
}
