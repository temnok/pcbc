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

func X2_I0603(topLabel, bottomLabel string) *eda.Component {
	return X2(topLabel, bottomLabel, smd.I0603)
}

func X2(topLabel, bottomLabel string, chip *eda.Component) *eda.Component {
	logo := boards.Logo.Arrange(transform.ScaleUniformly(0.7).Move(0.8, 1.35))
	//firm := boards.Firm.Arrange(transform.ScaleUniformly(0.5).Move(1.2, -1.4))
	rev := boards.Rev(2025, 5, 28).Arrange(transform.RotateDegrees(90).ScaleUniformly(0.5).Move(4.6, 0))

	header := greenconn.CSCC118(3, false, []string{topLabel, "GND", bottomLabel}).Arrange(transform.Move(-2, 0))
	pad := header.PadCenters()

	chip = chip.Arrange(transform.RotateDegrees(-90).Move(0.85, -0.45))
	pin := chip.PadCenters()

	mount := boards.MountHole.Arrange(transform.RotateDegrees(90).Move(3, 0))

	return &eda.Component{
		Cuts: path.Paths{
			path.RoundRect(10, 4, 0.5),
		},

		Components: eda.Components{
			header,
			chip,
			mount,

			logo,
			//firm,
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
