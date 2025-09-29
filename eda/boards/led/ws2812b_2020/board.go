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
	header = greenconn.CSCC118(3, false, []string{"DI", "VDD", "DO"}).Arrange(transform.Move(-3.2, 0))
	pads   = header.PadCenters()

	chip = worldsemi.WS2812B_2020.Arrange(transform.Move(0.6, 0))
	pins = chip.PadCenters()

	Board = &eda.Component{
		Cuts: path.Paths{
			path.RoundRect(11.5, 4, 1.1),
		},

		Tracks: path.Paths{
			eda.LinearTrack(pads[0], pins[3], 0, -1, 0),
			eda.LinearTrack(pads[1], pins[4], 0, -2.8, 0),
			eda.LinearTrack(pads[2], pins[1], 0, 1.1, 0.3, 0),
		},

		Nested: eda.Components{
			{
				ClearWidth:  eda.ClearOff,
				TracksWidth: 0.3,
				Tracks: path.Paths{
					eda.LinearTrack(pins[2], pins[2].Move(0.6, 0), -0.6),
				},
			},

			header,

			chip,

			boards.MountHole.Arrange(transform.RotateDegrees(45).Move(3.8, 0)),

			boards.Logo.Arrange(transform.ScaleUniformly(0.7).Move(5, 1.3)),

			boards.Rev(2025, 9, 27).Arrange(transform.ScaleUniformly(0.6).Move(4.5, -1.6)),

			{
				MarksWidth: 0.17,

				Nested: eda.Components{
					eda.CenteredText("WS2812").Arrange(transform.Scale(1.4, 0.7).Move(0.5, 1.55)),
					eda.CenteredText("B-2020").Arrange(transform.Scale(1.4, 0.7).Move(0.5, -1.55)),
				},
			},
		},
	}
)
