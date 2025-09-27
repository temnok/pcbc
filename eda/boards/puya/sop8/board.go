// Copyright © 2025 Alex Temnok. All rights reserved.

package sop8

import (
	"temnok/pcbc/eda"
	"temnok/pcbc/eda/boards"
	"temnok/pcbc/eda/lib/header/greenconn"
	"temnok/pcbc/eda/lib/pkg/sop"
	"temnok/pcbc/path"
	"temnok/pcbc/transform"
)

var (
	mount = boards.MountHole.Arrange(transform.RotateDegrees(90).Move(-5, 0))

	chip = sop.SOP8.Arrange(transform.Move(-0.6, 0))
	pin  = chip.PadCenters()

	header = greenconn.CSCC118(7, true, []string{"SWD", "PA2", "PA1", "VCC", "PA4", "PA3", "SWC"}).
		Arrange(transform.Move(4.4, 0))
	pad = header.PadCenters()

	Board = &eda.Component{
		Cuts: path.Paths{
			path.RoundRect(14, 8, 1),
		},

		Nested: eda.Components{
			mount,
			chip,
			header,

			eda.CenteredText("PY32").Arrange(transform.ScaleUniformly(1.3).Move(-5, 3.1)),
			eda.CenteredText("F002A").Arrange(transform.Scale(1.05, 1.3).Move(-5, 2.1)),
			boards.Logo.Arrange(transform.ScaleUniformly(1.3).Move(-5.8, -2.6)),
			boards.Firm.Arrange(transform.ScaleUniformly(0.5).Move(-4, -2.6)),
			boards.Rev(2025, 9, 27).Arrange(transform.ScaleUniformly(0.7).Move(2, -3.6)),

			{
				ClearOff: true,

				TracksWidth: 0.3,

				Tracks: eda.DeprecatedTracks(
					eda.DeprecatedTrack{pin[7]}.DY(-1).DY(2),
				),
			},
		},

		Tracks: eda.DeprecatedTracks(
			eda.DeprecatedTrack{pin[0]}.DY(0.8).YX(pad[3]),
			eda.DeprecatedTrack{pin[1]}.DY(0.8).YX(pad[4]),
			eda.DeprecatedTrack{pin[2]}.DY(0.8).DY(0.3).DX(2.2).YX(pad[5]),
			eda.DeprecatedTrack{pin[3]}.YX(pad[6]),
			eda.DeprecatedTrack{pin[4]}.YX(pad[0]),
			eda.DeprecatedTrack{pin[5]}.DY(-0.8).DY(-0.3).DX(2.2).YX(pad[1]),
			eda.DeprecatedTrack{pin[6]}.DY(-0.8).YX(pad[2]),
		),
	}
)
