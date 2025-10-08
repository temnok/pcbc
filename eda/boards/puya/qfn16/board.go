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
	mount = boards.MountHole.Arrange(transform.Move(0, -2.75))

	chip = qfn.QFN16G.Arrange(transform.RotateDegrees(-45).Move(0, 1.2))
	pin  = chip.PadCenters()

	leftLabels  = []string{"PB1", "A12", "SWD", "SWC", "PF2", "PA0", "PA1", "PA2", "GND"}
	rightLabels = []string{"PA8", "VCC", "PB0", "PA7", "PA6", "PA5", "PA4", "PA3", "GND"}

	header = &eda.Component{
		Nested: eda.Components{
			greenconn.CSCC118(9, true, leftLabels).Arrange(transform.Move(-4.5, 0)),
			greenconn.CSCC118(9, false, rightLabels).Arrange(transform.Move(4.5, 0)),
		},
	}
	pad = header.PadCenters()

	tracks = &eda.Component{
		Tracks: path.Paths{
			eda.LinearTrack(pin[0], pad[0]),
			eda.LinearTrack(pin[1], pad[1]),
			eda.LinearTrack(pin[2], pad[2], 0, 0, 0.4, -0.4, -1e-9),
			eda.LinearTrack(pin[3], pad[3]),
			eda.LinearTrack(pin[4], pad[4]),
			eda.LinearTrack(pin[5], pad[5]),
			eda.LinearTrack(pin[6], pad[6]),
			eda.LinearTrack(pin[7], pad[7]),
		},

		Nested: eda.Components{
			{
				ClearWidth:  eda.ClearOff,
				TracksWidth: 0.3,

				Tracks: path.Paths{
					eda.LinearTrack(pad[8], pad[8].Move(1.2, 0), -1.2),
				},
			},
		},
	}

	Board = &eda.Component{
		Cuts: path.Paths{
			path.RoundRect(14.3, 9.8, 1),
		},

		Nested: eda.Components{
			mount,
			chip,
			header,

			boards.Logo.Arrange(transform.ScaleUniformly(1.2).Move(-1.5, -1.1)),
			boards.Firm.Arrange(transform.ScaleUniformly(0.5).Move(1.35, -1.25)),
			eda.CenteredText("PY32").Arrange(transform.Scale(1.4, 1.6).Move(0, 4.1)),
			eda.CenteredText("F002A").Arrange(transform.Scale(1.2, 0.75).Move(0, -4.45)),

			boards.Rev(2025, 10, 7).Arrange(transform.ScaleUniformly(0.5).Move(6, -4.55)),

			tracks,
			tracks.Arrange(transform.MirrorX),
		},
	}
)
