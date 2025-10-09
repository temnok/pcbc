// Copyright © 2025 Alex Temnok. All rights reserved.

package ws2812b_2020

import (
	"temnok/pcbc/eda"
	"temnok/pcbc/eda/boards"
	"temnok/pcbc/eda/lib/header/greenconn"
	"temnok/pcbc/eda/lib/worldsemi"
	"temnok/pcbc/path"
	"temnok/pcbc/transform"
)

var (
	chip = worldsemi.WS2812B_2020.Arrange(transform.RotateDegrees(180).Move(-3.6, 0))
	pins = chip.PadCenters()

	header = greenconn.CSCC118(3, false, []string{"DO", "+V", "DI"}).Arrange(transform.Move(0.2, 0))
	pads   = header.PadCenters()

	Board = &eda.Component{
		Cuts: path.Paths{
			path.RoundRect(11.5, 4, 1.2),
		},

		Tracks: path.Paths{
			eda.LinearTrack(pads[0], pins[1], 0, 1, 0.3, -1e-9),
			eda.LinearTrack(pads[1], pins[4], 0, -1, 0),
			eda.LinearTrack(pads[2], pins[3], -1e-9, -1e-9),
		},

		Nested: eda.Components{
			{
				ClearWidth: eda.ClearOff,
				Tracks: path.Paths{
					eda.LinearTrack(pins[2], pins[2].Move(0.6, 0), -0.6),
				},
			},

			header,

			chip,

			boards.MountHole.Arrange(transform.Move(3.8, 0)),

			boards.Logo.Arrange(transform.ScaleUniformly(0.7).Move(2.8, 1.3)),

			boards.Rev(2025, 10, 8).Arrange(transform.ScaleUniformly(0.6).Move(3.8, -1.6)),

			{
				Transform: transform.Scale(0.9, 0.7).Move(-3.5, 0),

				MarksWidth: 0.17,

				Nested: eda.Components{
					eda.CenteredText("WS2812").Arrange(transform.Move(0, 2.1)),
					eda.CenteredText("B-2020").Arrange(transform.Move(0, -2.1)),
				},
			},

			boards.LogoBottom.Arrange(transform.ScaleUniformly(2)),
		},
	}
)
