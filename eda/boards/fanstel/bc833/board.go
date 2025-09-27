// Copyright © 2025 Alex Temnok. All rights reserved.

package bc833

import (
	"temnok/pcbc/eda"
	"temnok/pcbc/eda/boards"
	"temnok/pcbc/eda/lib/fanstel"
	"temnok/pcbc/eda/lib/header/greenconn"
	"temnok/pcbc/path"
	"temnok/pcbc/transform"
)

var (
	chip     = fanstel.BC833.Arrange(transform.Move(0, 4.25))
	chipPads = chip.PadCenters()

	leftConn = greenconn.CSCC118(12, false, []string{
		"P031", "P030", "P029", " VDD", "P003", "P002",
		"P028", " GND", "P020", "P017", "P004", "P005",
	}).Arrange(transform.Move(-8.2, 0))
	leftPads = leftConn.PadCenters()

	rightConn = greenconn.CSCC118(12, true, []string{
		"P010", "P009", " SWC", "SWD ", "P018", "P015",
		" D- ", " D+ ", "VBUS", "VDDH", "P011", "P109",
	}).Arrange(transform.Move(8.2, 0))
	rightPads = rightConn.PadCenters()

	Board = &eda.Component{
		TracksWidth: 0.25,
		ClearWidth:  0.25,

		Cuts: path.Paths{
			path.RoundRect(23, 14, 1.5),
		},

		Tracks: path.Paths{
			eda.LinearTrack(chipPads[1], leftPads[3]),
			eda.LinearTrack(chipPads[2], leftPads[4]),
			eda.LinearTrack(chipPads[3], leftPads[5]),
			eda.LinearTrack(chipPads[4], leftPads[6]),

			eda.LinearTrack(chipPads[6], leftPads[2], 0, 1, -0.7, -0.6, 0),
			eda.LinearTrack(chipPads[7], leftPads[1], 0, 0, 0.7, -1.45, -1.05, -1.3, 0),
			eda.LinearTrack(chipPads[8], leftPads[0], 0, 0, 1.4, -2, 0),

			eda.LinearTrack(chipPads[9], leftPads[10]),
			eda.LinearTrack(chipPads[10], leftPads[11]),

			eda.LinearTrack(chipPads[11], rightPads[11]),
			eda.LinearTrack(chipPads[12], rightPads[10]),
			eda.LinearTrack(chipPads[13], rightPads[9]),
			eda.LinearTrack(chipPads[14], rightPads[8]),

			eda.LinearTrack(chipPads[15], rightPads[6]),
			eda.LinearTrack(chipPads[16], rightPads[5]),
			eda.LinearTrack(chipPads[17], rightPads[4]),
			eda.LinearTrack(chipPads[18], rightPads[3]),
			eda.LinearTrack(chipPads[19], rightPads[7]),

			eda.LinearTrack(chipPads[20], leftPads[9], 2.3, 2, -1e-9),
			eda.LinearTrack(chipPads[21], leftPads[8], 2.7, 2, -1e-9),

			eda.LinearTrack(chipPads[22], rightPads[2], 0, -1e-9, -0.7, 0.6, 0),
			eda.LinearTrack(chipPads[23], rightPads[1], 1.2, 1.5, 0),
			eda.LinearTrack(chipPads[24], rightPads[0]),
		},

		Nested: eda.Components{
			chip,
			leftConn,
			rightConn,

			boards.Logo.Arrange(transform.Scale(2, 2).Move(-3.6, -4.6)),
			eda.CenteredText("BC833").Arrange(transform.Scale(1.1, 2.5).Move(3.4, -4)),
			eda.CenteredText("nRF52833").Arrange(transform.Scale(0.85, 2).Move(3.15, -5.9)),

			{
				ClearOff:    true,
				TracksWidth: 0.35,

				Tracks: path.Paths{
					eda.LinearTrack(chipPads[5], chipPads[5].Move(1, 0)),
					eda.LinearTrack(chipPads[5], leftPads[7]),
					eda.LinearTrack(leftPads[7], leftPads[7].Move(-2, 0)),

					eda.LinearTrack(chipPads[26], chipPads[26].Move(0, 2)),
					eda.LinearTrack(chipPads[26], chipPads[26].Move(-2, 0)),
					eda.LinearTrack(chipPads[26], chipPads[26].Move(-2, 2)),
				},
			},

			boards.MountHole.Arrange(transform.Move(0, -4.7)),
		},
	}
)
