// Copyright © 2025 Alex Temnok. All rights reserved.

package sop8

import (
	"temnok/pcbc/eda"
	"temnok/pcbc/eda/lib/header/greenconn"
	"temnok/pcbc/eda/lib/pkg/sop"
	"temnok/pcbc/eda/pcbc"
	"temnok/pcbc/path"
	"temnok/pcbc/transform"
)

var (
	mount = pcbc.MountHole.Arrange(transform.Rotate(90).Move(-5, 0))

	chip = sop.SOP8.Arrange(transform.Move(-0.6, 0))
	pin  = chip.PadCenters()

	header = greenconn.CSCC118(7, true, []string{"SWD", "PA2", "PA1", "VCC", "PA4", "PA3", "SWC"}).
		Arrange(transform.Move(4.4, 0))
	pad = header.PadCenters()

	Board = &eda.Component{
		Cuts: path.Paths{
			path.RoundRect(14, 8, 1),
		},

		Components: eda.Components{
			mount,
			chip,
			header,

			eda.CenteredText("PY32").Arrange(transform.Scale(1.3, 2.8).Move(-5, 2.5)),
			eda.CenteredText("F002A").Arrange(transform.Scale(1.1, 2.8).Move(-5, -2.5)),
		},

		Tracks: eda.Tracks(
			eda.Track{pin[0]}.DY(0.8).YX(pad[3]),
			eda.Track{pin[1]}.DY(0.8).YX(pad[4]),
			eda.Track{pin[2]}.DY(0.8).DY(0.3).DX(2.2).YX(pad[5]),
			eda.Track{pin[3]}.YX(pad[6]),
			eda.Track{pin[4]}.YX(pad[0]),
			eda.Track{pin[5]}.DY(-0.8).DY(-0.3).DX(2.2).YX(pad[1]),
			eda.Track{pin[6]}.DY(-0.8).YX(pad[2]),
		),

		GroundTracks: eda.Tracks(
			eda.Track{pin[7]}.DY(-1).DY(2),
		),
	}
)
