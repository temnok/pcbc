// Copyright © 2025 Alex Temnok. All rights reserved.

package sop8

import (
	"github.com/temnok/pcbc/eda"
	"github.com/temnok/pcbc/eda/boards"
	"github.com/temnok/pcbc/eda/lib/header/greenconn"
	"github.com/temnok/pcbc/eda/lib/pkg/sop"
	"github.com/temnok/pcbc/path"
	"github.com/temnok/pcbc/transform"
)

var (
	mount = boards.MountHole.Arrange(transform.Move(-5, 0))

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

			{
				MarksWidth: 0.15,
				Nested: eda.Components{
					eda.CenteredText("PY32").Arrange(transform.ScaleUniformly(1.3).Move(-5, 3.1)),
					eda.CenteredText("F002A").Arrange(transform.Scale(1.05, 1.3).Move(-5, 2.0)),
				},
			},

			boards.Logo.Arrange(transform.ScaleUniformly(1.3).Move(-5.8, -2.6)),
			boards.Firm.Arrange(transform.ScaleUniformly(0.5).Move(-4, -2.6)),
			boards.Rev(2025, 10, 7).Arrange(transform.ScaleUniformly(0.7).Move(2, -3.6)),

			{
				ClearWidth:  eda.ClearOff,
				TracksWidth: 0.3,

				Tracks: path.Paths{
					eda.LinearTrack(pin[7], pin[7].Move(0, 1), -1),
				},
			},

			boards.LogoBottom.Arrange(transform.ScaleUniformly(4)),
		},

		Tracks: path.Paths{
			eda.LinearTrack(pin[0], pad[3]),
			eda.LinearTrack(pin[1], pad[4]),
			eda.LinearTrack(pin[2], pad[5], 0.9, 0.7),
			eda.LinearTrack(pin[3], pad[6]),
			eda.LinearTrack(pin[4], pad[0]),
			eda.LinearTrack(pin[5], pad[1], 0.9, -0.7),
			eda.LinearTrack(pin[6], pad[2]),
		},
	}
)
