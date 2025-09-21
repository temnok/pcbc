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

var hole = path.Circle(1.45)

var boardTop = &eda.Component{
	Layer: 1,

	Tracks: eda.Tracks(
		eda.Track{{-7.5, 2.75}}.DY(-5.5).DX(3.5).DY(-3),
		eda.Track{{7.5, 2.75}}.DY(-5.5).DX(-3.5).DY(-3),
		eda.Track{{-6, -6.2}, {6, -6.2}},
	),

	Nested: eda.Components{
		{
			Transform: transform.Move(0, -5.5),

			Nested: eda.Components{
				greenconn.CSCC118(13, false,
					[]string{"3V7", "3V7", "3V7", "3V7", "3V7", "3V7",
						"3V7", "3V7", "3V7", "3V7", "3V7", "3V7", "3V7"},
				).Arrange(transform.RotateDegrees(90).Move(0, -0.7)),
			},
		},

		holder.LIR1254.Arrange(transform.Move(0, 2.75)),

		boards.Logo.Arrange(transform.Scale(1.6, 1.6).Move(-8.4, -5.3)),

		boards.Firm.Arrange(transform.Scale(0.8, 0.8).Move(8.4, -5.3)),

		boards.Rev(2025, 9, 21).Arrange(transform.RotateDegrees(90).Scale(0.8, 0.8).Move(9.2, -0.4)),

		eda.CenteredText("LIR1254").Arrange(transform.Scale(1, 1.6).Move(-7.5, 6)),

		eda.CenteredText("3.7V").Arrange(transform.Scale(1.5, 1.5).Move(7.8, 6)),
	},
}

var Board = &eda.Component{
	TracksWidth: 0.55, // more power!
	ClearWidth:  0.3,

	Nested: eda.Components{
		{
			CutsOuter: true,

			Cuts: path.Paths{
				hole.Transform(transform.Move(-7.5, -7.5)),
				hole.Transform(transform.Move(-7.5, 7.5)),
				hole.Transform(transform.Move(7.5, -7.5)),
				hole.Transform(transform.Move(7.5, 7.5)),
			},
		},

		{
			Cuts: path.Paths{
				path.RoundRect(20, 18, 2),
			},
		},

		boardTop,

		boards.MountHole15(1).Arrange(transform.Move(-7.5, -2.5)),

		boards.MountHole15(1).Arrange(transform.Move(7.5, -2.5)),
	},
}
