// Copyright © 2025 Alex Temnok. All rights reserved.

package lir1254

import (
	"temnok/pcbc/eda"
	"temnok/pcbc/eda/boards"
	"temnok/pcbc/eda/lib/battery/holder"
	"temnok/pcbc/eda/lib/header/mph100imp40f"
	"temnok/pcbc/path"
	"temnok/pcbc/transform"
)

var Board = &eda.Component{
	TrackWidth: 0.5, // more power!

	Cuts: path.Paths{
		path.RoundRect(21, 14, 1),
	},

	Inner: eda.Components{
		{
			Transform: transform.Move(0, -5.5),

			Inner: eda.Components{
				mph100imp40f.G_V_SP(8),

				eda.CenteredTextRow(2.54, "3V7", "3V7", "3V7", "3V7", "3V7", "3V7", "3V7", "3V7").
					Arrange(transform.Move(0, 1.8)),
			},
		},

		holder.LIR1254.Arrange(transform.Move(0, 2.75)),

		boards.MountHole.Arrange(transform.RotateDegrees(-45).Move(-7.5, -1.75)),

		boards.MountHole.Arrange(transform.RotateDegrees(45).Move(7.5, -1.75)),

		boards.Logo.Arrange(transform.Scale(1.2, 1.2).Move(-5, -2)),

		eda.CenteredText("LIR1254").Arrange(transform.Scale(1, 1.5).Move(-7.8, 6)),

		eda.CenteredText("3.7V").Arrange(transform.Scale(1.5, 1.5).Move(7.8, 6)),
	},
}

func init() {
	pad := Board.PadCenters()

	Board.Tracks = eda.Tracks(
		eda.Track{pad[0]}.XY(pad[7]),
		eda.Track{pad[0]}.DX(-0.8).YX(pad[8]),
		eda.Track{pad[7]}.DX(0.8).YX(pad[9]),
	)
}
