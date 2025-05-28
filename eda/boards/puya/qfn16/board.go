// Copyright © 2025 Alex Temnok. All rights reserved.

package qfn16

import (
	"temnok/pcbc/eda"
	"temnok/pcbc/eda/boards"
	"temnok/pcbc/eda/lib/header/greenconn"
	"temnok/pcbc/eda/lib/pkg/qfn"
	"temnok/pcbc/path"
	"temnok/pcbc/transform"
)

var (
	mount    = boards.MountHole.Arrange(transform.RotateDegrees(90).Move(0, -2.8))
	mountPad = mount.PadCenters()

	chip = qfn.QFN16G.Arrange(transform.RotateDegrees(-45).Move(0, 1.2))
	pin  = chip.PadCenters()

	leftLabels  = []string{"PB1", "A12", "SWD", "SWC", "PF2", "PA0", "PA1", "PA2", "GND"}
	rightLabels = []string{"PA8", "VCC", "PB0", "PA7", "PA6", "PA5", "PA4", "PA3", "GND"}

	header = &eda.Component{
		Components: eda.Components{
			greenconn.CSCC118(9, true, leftLabels).Arrange(transform.Move(-4.5, 0)),
			greenconn.CSCC118(9, false, rightLabels).Arrange(transform.Move(4.5, 0)),
		},
	}
	pad = header.PadCenters()

	Board = &eda.Component{
		Cuts: path.Paths{
			path.RoundRect(14.3, 9.8, 1),
		},

		Components: eda.Components{
			mount,
			chip,
			header,

			boards.Logo.Arrange(transform.ScaleUniformly(1.2).Move(-1.5, -1.1)),
			boards.Firm.Arrange(transform.ScaleUniformly(0.5).Move(1.35, -1.25)),
			eda.CenteredText("PY32").Arrange(transform.Scale(1.3, 1.5).Move(0, 4.1)),
			eda.CenteredText("F002A").Arrange(transform.Scale(1.0, 0.9).Move(0, -4.4)),

			boards.Rev(2025, 5, 28).Arrange(transform.ScaleUniformly(0.5).Move(6, -4.55)),

			{
				NoClear: true,

				Tracks: eda.Tracks(
					eda.Track{mountPad[2]}.YX(pad[8]).DX(-1.3),
					eda.Track{mountPad[4]}.YX(pad[17]).DX(1.3),
				),
			},
		},

		Tracks: eda.Tracks(
			eda.Track{pin[0]}.YX(pad[0]),
			eda.Track{pin[1]}.YX(pad[1]),
			eda.Track{pin[2]}.DY(0.3).DX(-1.3).YX(pad[2]),
			eda.Track{pin[3]}.YX(pad[3]),
			eda.Track{pin[4]}.YX(pad[4]),
			eda.Track{pin[5]}.YX(pad[5]),
			eda.Track{pin[6]}.YX(pad[6]),
			eda.Track{pin[7]}.YX(pad[7]),

			eda.Track{pin[8]}.YX(pad[16]),
			eda.Track{pin[9]}.YX(pad[15]),
			eda.Track{pin[10]}.YX(pad[14]),
			eda.Track{pin[11]}.YX(pad[13]),
			eda.Track{pin[12]}.YX(pad[12]),
			eda.Track{pin[13]}.DY(0.3).DX(1.3).YX(pad[11]),
			eda.Track{pin[14]}.YX(pad[10]),
			eda.Track{pin[15]}.YX(pad[9]),
		),
	}
)
