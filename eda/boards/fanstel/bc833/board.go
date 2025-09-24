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
	chip     = fanstel.BC833.Arrange(transform.Move(0, 4.3))
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

		Tracks: eda.Tracks(
			eda.TrackV2(chipPads[1], leftPads[3]),
			eda.TrackV2(chipPads[2], leftPads[4]),
			eda.TrackV2(chipPads[3], leftPads[5]),
			eda.TrackV2(chipPads[4], leftPads[6]),

			eda.TrackV2(chipPads[6], leftPads[2], 1, -0.7, -0.6, 0),
			eda.TrackV2(chipPads[7], leftPads[1], 0, 0.7, -1.45, -1.1, -1.3, 0),
			eda.TrackV2(chipPads[8], leftPads[0], 0, 1.4, -2, 0),

			eda.TrackV2(chipPads[9], leftPads[10]),
			eda.TrackV2(chipPads[10], leftPads[11]),

			eda.TrackV2(chipPads[11], rightPads[11]),
			eda.TrackV2(chipPads[12], rightPads[10]),
			eda.TrackV2(chipPads[13], rightPads[9]),
			eda.TrackV2(chipPads[14], rightPads[8]),

			eda.TrackV2(chipPads[15], rightPads[6]),
			eda.TrackV2(chipPads[16], rightPads[5]),
			eda.TrackV2(chipPads[17], rightPads[4]),
			eda.TrackV2(chipPads[18], rightPads[3]),
			eda.TrackV2(chipPads[19], rightPads[7]),

			eda.TrackV2(chipPads[20], leftPads[9], -0.01, 2.3, 2, -0.01),
			eda.TrackV2(chipPads[21], leftPads[8], -0.01, 2.7, 2, -0.01),

			eda.TrackV2(chipPads[22], rightPads[2], -0.01, -0.7, 0.6, 0),
			eda.TrackV2(chipPads[23], rightPads[1], -0.01, 1.25, 1.5, 0),
			eda.TrackV2(chipPads[24], rightPads[0]),
		),

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

				Tracks: eda.Tracks(
					eda.TrackV2(chipPads[5], chipPads[5].Move(1, 0)),
					eda.TrackV2(chipPads[5], leftPads[7]),
					eda.TrackV2(leftPads[7], leftPads[7].Move(-2, 0)),

					eda.TrackV2(chipPads[26], chipPads[26].Move(0, 2)),
					eda.TrackV2(chipPads[26], chipPads[26].Move(-2, 0)),
					eda.TrackV2(chipPads[26], chipPads[26].Move(-2, 2)),
				),
			},

			boards.MountHole.Arrange(transform.Move(0, -4.7)),
		},
	}
)
