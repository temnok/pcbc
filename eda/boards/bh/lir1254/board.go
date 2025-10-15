// Copyright © 2025 Alex Temnok. All rights reserved.

package lir1254

import (
	"github.com/temnok/pcbc/eda"
	"github.com/temnok/pcbc/eda/boards"
	"github.com/temnok/pcbc/eda/lib/battery/holder"
	"github.com/temnok/pcbc/eda/lib/header/greenconn"
	"github.com/temnok/pcbc/path"
	"github.com/temnok/pcbc/transform"
	"github.com/temnok/pcbc/util"
)

var (
	conn = greenconn.CSCC118(11, true, append([]string{""}, util.Repeat("3V7", 9)...)).
		Arrange(transform.RotateDegrees(90).Move(0, -5.2))

	connPads = conn.PadCenters()

	hold     = holder.LIR1254.Arrange(transform.Move(0, 4.2))
	holdPads = hold.PadCenters()

	Board = &eda.Component{
		TracksWidth: 0.5,

		Nested: eda.Components{
			(&eda.Component{
				Tracks: path.Paths{
					eda.LinearTrack(holdPads[0], connPads[5].Move(0, -1), 0, 0, 0.9, -7, -1e-9),
				},
			}).Clone(2, transform.MirrorX),

			{
				Cuts: path.Paths{
					path.RoundRect(18, 16, 3),
				},
			},

			boards.MountHole.CloneX(2, 12).Arrange(transform.Move(0, -2.4)),

			conn,

			hold,

			boards.Logo.Arrange(transform.ScaleUniformly(1.8).Move(-7, 0.3)),

			boards.Firm.Arrange(transform.ScaleUniformly(0.9).Move(7, 0.3)),

			boards.Rev(2025, 10, 7).Arrange(transform.RotateDegrees(90).Scale(0.8, 0.8).Move(8.35, -5)),

			{
				MarksWidth: 0.18,

				Nested: eda.Components{
					eda.CenteredText("LIR1254").Arrange(transform.Scale(2.1, 1.3).Move(0, -2.3)),
				},
			},

			boards.LogoBottom.Arrange(transform.ScaleUniformly(7)),
		},
	}
)
