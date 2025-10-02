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
		Arrange(transform.RotateDegrees(90).Move(0, -4.5))

	connPads = conn.PadCenters()

	hold     = holder.LIR1254.Arrange(transform.Move(0, 4))
	holdPads = hold.PadCenters()

	Board = &eda.Component{
		TracksWidth: 0.5,
		ClearWidth:  0.3,

		Tracks: path.Paths{
			eda.LinearTrack(holdPads[0], connPads[3], 0, 2, -1e-9),
			eda.LinearTrack(holdPads[1], connPads[7], -1e-9, -2, 0),
			eda.LinearTrack(connPads[0].Move(0, 1), connPads[10].Move(0, 1)),
		},

		Nested: eda.Components{
			{
				Cuts: path.Paths{
					path.RoundRect(18, 15, 1.4),
				},
			},

			boards.MountHoleV2.Arrange(transform.Move(0, -2)).CloneX(2, 12),

			conn,

			hold,

			boards.Logo.Arrange(transform.Scale(1.6, 1.6).Move(-7.5, 0.3)),

			boards.Firm.Arrange(transform.Scale(0.8, 0.8).Move(7.5, 0.3)),

			boards.Rev(2025, 10, 1).Arrange(transform.RotateDegrees(90).Scale(0.8, 0.8).Move(8.35, -2)),

			eda.CenteredText("LIR").Arrange(transform.Scale(1.5, 2).Move(-7.3, -4.5)),
			eda.CenteredText("1254").Arrange(transform.Scale(1.1, 2).Move(-7.3, -6)),

			eda.CenteredText("3.7").Arrange(transform.Scale(1.5, 2).Move(7.3, -4.5)),
			eda.CenteredText("V").Arrange(transform.Scale(2, 2).Move(7.3, -6)),
		},
	}
)
