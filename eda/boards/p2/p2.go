// Copyright © 2025 Alex Temnok. All rights reserved.

package p2

import (
	"temnok/pcbc/eda"
	"temnok/pcbc/eda/boards"
	"temnok/pcbc/eda/lib/header/greenconn"
	"temnok/pcbc/eda/lib/pkg/smd"
	"temnok/pcbc/path"
	"temnok/pcbc/transform"
)

func P2_I0402(topLabel, bottomLabel string) *eda.Component {
	return P2(topLabel, bottomLabel, smd.I0402)
}

func P2(topLabel, bottomLabel string, chip *eda.Component) *eda.Component {
	chip = chip.Arrange(transform.RotateDegrees(-90).Move(-3.5, 0))
	pin := chip.PadCenters()

	header := greenconn.CSCC118(2, false, []string{topLabel, bottomLabel}).Arrange(transform.Move(-0.8, 0))
	pad := header.PadCenters()

	mount := boards.MountHoleV2.Arrange(transform.Move(2.9, 0))

	return &eda.Component{
		Cuts: path.Paths{
			path.RoundRect(9, 3, 1),
		},

		Tracks: path.Paths{
			eda.LinearTrack(pad[0], pin[0], 0, 0),
			eda.LinearTrack(pad[1], pin[1], 0, 0),
		},

		Nested: eda.Components{
			header,
			chip,
			mount,

			boards.Logo.Arrange(transform.ScaleUniformly(0.6).Move(1.5, 1)),
			boards.Rev(2025, 10, 4).Arrange(transform.ScaleUniformly(0.5).Move(-2.3, -1.1)),
		},
	}
}
