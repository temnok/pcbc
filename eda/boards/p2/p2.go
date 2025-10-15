// Copyright © 2025 Alex Temnok. All rights reserved.

package p2

import (
	"github.com/temnok/pcbc/eda"
	"github.com/temnok/pcbc/eda/boards"
	"github.com/temnok/pcbc/eda/lib/header/greenconn"
	"github.com/temnok/pcbc/eda/lib/pkg/smd"
	"github.com/temnok/pcbc/path"
	"github.com/temnok/pcbc/transform"
)

func P2_I0402(topLabel, bottomLabel string) *eda.Component {
	return P2(topLabel, bottomLabel, smd.I0402)
}

func P2(topLabel, bottomLabel string, chip *eda.Component) *eda.Component {
	chip = chip.Arrange(transform.RotateDegrees(-90).Move(-3.8, 0))
	pin := chip.PadCenters()

	header := greenconn.CSCC118(3, false, []string{topLabel, "GND", bottomLabel}).Arrange(transform.Move(-0.8, 0))
	pad := header.PadCenters()

	mount := boards.MountHole.Arrange(transform.Move(3, 0))

	return &eda.Component{
		Cuts: path.Paths{
			path.RoundRect(9.8, 3.8, 1),
		},

		Tracks: path.Paths{
			eda.LinearTrack(pad[0], pin[0], 0, 0),
			eda.LinearTrack(pad[2], pin[1], 0, 0),
		},

		Nested: eda.Components{
			header,
			chip,
			mount,

			{
				ClearWidth: eda.ClearOff,

				Tracks: path.Paths{
					eda.LinearTrack(pad[1], pad[1].Move(1.2, 0), -1.2),
				},
			},

			boards.Logo.Arrange(transform.ScaleUniformly(0.7).Move(1.8, 1.2)),
			boards.Rev(2025, 10, 8).Arrange(transform.ScaleUniformly(0.5).Move(3, -1.5)),

			boards.LogoBottom.Arrange(transform.ScaleUniformly(2)),
		},
	}
}
