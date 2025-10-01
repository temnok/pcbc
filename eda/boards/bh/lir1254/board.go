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
	conn = greenconn.CSCC118(19, false, util.Repeat("3V7", 19)).
		Arrange(transform.RotateDegrees(90).Move(0, -5.2))

	connPads = conn.PadCenters()

	hold     = holder.LIR1254.Arrange(transform.Move(0, 3.75))
	holdPads = hold.PadCenters()

	Board = &eda.Component{
		TracksWidth: 0.55, // more power!
		ClearWidth:  0.3,

		Tracks: path.Paths{
			eda.LinearTrack(holdPads[0], connPads[6], 0, 2, -1e-9),
			eda.LinearTrack(holdPads[1], connPads[12], -1e-9, -2, 0),
			eda.LinearTrack(connPads[0].Move(0, -1), connPads[18].Move(0, -1)),
		},

		Nested: eda.Components{
			{
				Cuts: path.Paths{
					path.RoundRect(21, 16, 1.4),
				},
			},

			boards.MountHole.CloneX(2, 16).Arrange(transform.Move(0, -1)),

			conn,

			hold,

			boards.Logo.Arrange(transform.Scale(1.6, 1.6).Move(-5.2, -1.3)),

			boards.Firm.Arrange(transform.Scale(0.8, 0.8).Move(5, -2)),

			boards.Rev(2025, 9, 27).Arrange(transform.RotateDegrees(90).Scale(0.8, 0.8).Move(9.2, 0.6)),

			eda.CenteredText("LIR1254").Arrange(transform.Scale(1, 1.6).Move(-7.5, 7)),

			eda.CenteredText("3.7V").Arrange(transform.Scale(1.5, 1.5).Move(7.8, 7)),
		},
	}
)
