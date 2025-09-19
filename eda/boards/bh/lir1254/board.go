// Copyright © 2025 Alex Temnok. All rights reserved.

package lir1254

import (
	"temnok/pcbc/eda"
	"temnok/pcbc/eda/boards"
	"temnok/pcbc/eda/lib/battery/holder"
	"temnok/pcbc/eda/lib/header/greenconn"
	"temnok/pcbc/path"
	"temnok/pcbc/transform"
)

var Board = &eda.Component{
	TracksWidth: 0.5, // more power!

	Cuts: path.Paths{
		path.RoundRect(20, 16, 1).Transform(transform.Move(0, -1)),
	},

	Inner: eda.Components{
		{
			Transform: transform.Move(0, -5.5),

			Inner: eda.Components{
				greenconn.CSCC118(15, true,
					[]string{"3V7", "3V7", "3V7", "3V7", "3V7", "3V7", "3V7",
						"3V7", "3V7", "3V7", "3V7", "3V7", "3V7", "3V7", "3V7"},
				).Arrange(transform.RotateDegrees(90).Move(0, -0.7)),
			},
		},

		holder.LIR1254.Arrange(transform.Move(0, 2.75)),

		boards.MountHole15.Arrange(transform.Move(-7.5, -2.5)),

		boards.MountHole15.Arrange(transform.Move(7.5, -2.5)),

		boards.Logo.Arrange(transform.Scale(1.6, 1.6).Move(-4.8, -2.5)),

		boards.Firm.Arrange(transform.Scale(0.8, 0.8).Move(4.8, -2.7)),

		eda.CenteredText("LIR1254").Arrange(transform.Scale(1, 1.6).Move(-7.5, 6)),

		eda.CenteredText("3.7V").Arrange(transform.Scale(1.5, 1.5).Move(7.8, 6)),
	},
}

func init() {
	pad := Board.PadCenters()

	track := eda.Track{pad[15]}.DY(-5.5).DX(3.5).DY(-3)

	Board.Tracks = eda.Tracks(
		track,
		track.Apply(transform.MirrorX()),
		eda.Track{{-7, -6.2}, {7, -6.2}},
	)
}
