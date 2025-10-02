// Copyright © 2025 Alex Temnok. All rights reserved.

package lir1254

import (
	"temnok/pcbc/eda"
	"temnok/pcbc/eda/boards"
	"temnok/pcbc/eda/lib/battery/holder"
	"temnok/pcbc/eda/lib/header/greenconn"
	"temnok/pcbc/path"
	"temnok/pcbc/transform"
	"temnok/pcbc/util"
)

var (
	conn = greenconn.CSCC118(11, true, append([]string{""}, util.Repeat("3V7", 9)...)).
		Arrange(transform.RotateDegrees(90).Move(0, -4.4))

	connPads = conn.PadCenters()

	hold     = holder.LIR1254.Arrange(transform.Move(0, 4))
	holdPads = hold.PadCenters()

	Board = &eda.Component{
		TracksWidth: 0.5,
		ClearWidth:  0.25,

		Nested: eda.Components{
			(&eda.Component{
				Tracks: path.Paths{
					eda.LinearTrack(holdPads[0], connPads[5].Move(0, -1), 0, 0, 0.9, -6, -1e-9),
				},
			}).Clone(2, transform.MirrorX),

			{
				Cuts: path.Paths{
					path.RoundRect(18, 14.5, 2),
				},
			},

			boards.MountHoleV2.Arrange(transform.Move(0, -1.5)).CloneX(2, 12),

			conn,

			hold,

			boards.Logo.Arrange(transform.Scale(1.6, 1.6).Move(-7.5, 0.5)),

			boards.Firm.Arrange(transform.Scale(0.8, 0.8).Move(7.5, 0.5)),

			boards.Rev(2025, 10, 1).Arrange(transform.RotateDegrees(90).Scale(0.8, 0.8).Move(8.35, -1.5)),

			eda.CenteredText("LIR").Arrange(transform.Scale(1.5, 2).Move(-7.3, -4)),
			eda.CenteredText("1254").Arrange(transform.Scale(1.1, 2).Move(-7.3, -5.5)),

			eda.CenteredText("3.7").Arrange(transform.Scale(1.5, 2).Move(7.3, -4)),
			eda.CenteredText("V").Arrange(transform.Scale(2, 2).Move(7.3, -5.5)),
		},
	}
)
