// Copyright © 2025 Alex Temnok. All rights reserved.

package x2

import (
	"temnok/pcbc/eda"
	"temnok/pcbc/eda/boards"
	"temnok/pcbc/eda/lib/header/greenconn"
	"temnok/pcbc/eda/lib/pkg/smd"
	"temnok/pcbc/path"
	"temnok/pcbc/transform"
)

func X2_I0402(topLabel, bottomLabel string) *eda.Component {
	return X2(topLabel, bottomLabel, smd.I0402)
}

func X2(topLabel, bottomLabel string, chip *eda.Component) *eda.Component {
	logo := boards.Logo.Arrange(transform.ScaleUniformly(0.9).Move(0.8, 1.1))
	rev := boards.Rev(2025, 9, 27).Arrange(transform.ScaleUniformly(0.5).Move(4.15, -1.5))

	header := greenconn.CSCC118(3, false, []string{topLabel, "GND", bottomLabel}).Arrange(transform.Move(-2, 0))
	pad := header.PadCenters()

	chip = chip.Arrange(transform.RotateDegrees(-90).Move(0.85, -0.65))
	pin := chip.PadCenters()

	mount := boards.MountHole.Arrange(transform.RotateDegrees(45).Move(3, 0))

	return &eda.Component{
		Cuts: path.Paths{
			path.RoundRect(9.75, 3.75, 0.75),
		},

		Nested: eda.Components{
			header,
			chip,
			mount,

			logo,
			rev,

			{
				ClearOff: true,

				TracksWidth: 0.3,

				Tracks: path.Paths{
					eda.LinearTrack(pad[1], pad[1].Move(1.2, 0), -1.2),
				},
			},
		},

		Tracks: path.Paths{
			eda.LinearTrack(pad[0], pin[0]),
			eda.LinearTrack(pad[2], pin[1]),
		},
	}
}
