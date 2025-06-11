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
	rev := boards.Rev(2025, 6, 7).Arrange(transform.ScaleUniformly(0.5).Move(4.15, -1.5))

	header := greenconn.CSCC118(3, false, []string{topLabel, "GND", bottomLabel}).Arrange(transform.Move(-2, 0))
	pad := header.PadCenters()

	chip = chip.Arrange(transform.RotateDegrees(-90).Move(0.85, -0.65))
	pin := chip.PadCenters()

	mount := boards.MountHole.Arrange(transform.RotateDegrees(45).Move(3, 0))

	return &eda.Component{
		Cuts: path.Paths{
			path.RoundRect(9.75, 3.75, 0.5),
		},

		Components: eda.Components{
			header,
			chip,
			mount,

			logo,
			rev,

			{
				NoClear: true,

				Tracks: eda.Tracks(
					eda.Track{pad[1]}.DX(-1.2),
					eda.Track{pad[1]}.DX(1.2),
				),
			},
		},

		Tracks: eda.Tracks(
			eda.Track{pad[0]}.XY(pin[0]),
			eda.Track{pad[2]}.XY(pin[1]),
		),
	}
}
