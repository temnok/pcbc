// Copyright © 2025 Alex Temnok. All rights reserved.

package ts026a

import (
	"github.com/temnok/pcbc/eda"
	"github.com/temnok/pcbc/eda/boards"
	"github.com/temnok/pcbc/eda/lib/header/greenconn"
	"github.com/temnok/pcbc/eda/lib/ts/hyp"
	"github.com/temnok/pcbc/path"
	"github.com/temnok/pcbc/transform"
)

var (
	chip = hyp.Switch1TS026A.Arrange(transform.RotateDegrees(90).Move(-3.6, 0))
	pin  = chip.PadCenters()

	header = greenconn.CSCC118(3, false, []string{"TS", "GND", ""}).Arrange(transform.Move(-0.4, 0))
	pad    = header.PadCenters()

	Board = &eda.Component{
		Cuts: path.Paths{
			path.RoundRect(10.2, 4.2, 1),
		},

		Tracks: path.Paths{
			eda.LinearTrack(pad[0], pin[0], 0, 0),
			eda.LinearTrack(pad[2], pin[2], 0, 0),
		},

		Nested: eda.Components{
			header,
			chip,

			boards.MountHole.Arrange(transform.Move(3.2, 0)),

			{
				ClearWidth: eda.ClearOff,

				Tracks: path.Paths{
					eda.LinearTrack(pad[1], pad[1].Move(1.5, 0), -1.5),
				},
			},

			boards.Logo.Arrange(transform.ScaleUniformly(0.7).Move(2.1, 1.4)),
			boards.Rev(2025, 10, 8).Arrange(transform.ScaleUniformly(0.5).Move(3.2, -1.7)),
			boards.LogoBottom.Arrange(transform.ScaleUniformly(2.3)),
		},
	}
)
