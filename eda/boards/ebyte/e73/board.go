// Copyright © 2025 Alex Temnok. All rights reserved.

package e73

import (
	"temnok/pcbc/eda"
	"temnok/pcbc/eda/boards"
	"temnok/pcbc/eda/lib/ebyte"
	"temnok/pcbc/eda/lib/header/greenconn"
	"temnok/pcbc/path"
	"temnok/pcbc/transform"
)

var (
	e73      = ebyte.E73.Arrange(transform.Move(0, 2.67))
	chipPads = e73.PadCenters()

	leftConn = greenconn.CSCC118(13, false, []string{
		"P111", "P110", "P003", " AI4", "P113", " AI0",
		"AI5 ", " AI7", "AI6 ", " XL2", "XL1 ", " AI3", "P109",
	}).Arrange(transform.Move(-11.75, 1.7))
	leftPads = leftConn.PadCenters()

	rightConn = greenconn.CSCC118(13, true, []string{
		" NF2", "NF1 ", " SWC", "SWD ", "P024", "P013", "  D+",
		"D-  ", "VBUS", "RST ", "DCCH", "VDDH", " VDD",
	}).Arrange(transform.Move(11.75, 1.7))
	rightPads = rightConn.PadCenters()

	leftTracks = path.Paths{
		eda.LinearTrack(chipPads[0], leftPads[0]),
		eda.LinearTrack(chipPads[1], leftPads[1]),
		eda.LinearTrack(chipPads[2], leftPads[2], 0, -0.5, 0),
		eda.LinearTrack(chipPads[3], leftPads[3], 0, -0.5, 0),
		eda.LinearTrack(chipPads[5], leftPads[4], 0, -0.5, 0),
		eda.LinearTrack(chipPads[6], leftPads[5], 0, -0.5, 0),
		eda.LinearTrack(chipPads[7], leftPads[6], 0, -0.5, 0),
		eda.LinearTrack(chipPads[8], leftPads[7], 0, -0.5, 0),
		eda.LinearTrack(chipPads[9], leftPads[8], 0, -0.5, 0),

		eda.LinearTrack(chipPads[10], leftPads[10], 0, -1.9, 0),
		eda.LinearTrack(chipPads[11], leftPads[9], 0, 0.8, -0.5, -3.2, 2, 0.5, -1e-9),
		eda.LinearTrack(chipPads[12], leftPads[11], 0, -1e-9, -1e-9, -0.7, 0.5, 3.5, 0),
		eda.LinearTrack(chipPads[13], leftPads[12], 0, -1e-9, -1e-9, -0.7, 1, 4.5, 0),
	}

	rightTracks = path.Paths{
		eda.LinearTrack(chipPads[27], rightPads[0]),
		eda.LinearTrack(chipPads[26], rightPads[1]),
		eda.LinearTrack(chipPads[25], rightPads[2], 0, 0.5, 0),
		eda.LinearTrack(chipPads[24], rightPads[3], 0, 0.5, 0),
		eda.LinearTrack(chipPads[23], rightPads[4], 0, 0.5, 0),
		eda.LinearTrack(chipPads[22], rightPads[5], 0, 0.5, 0),
		eda.LinearTrack(chipPads[21], rightPads[6], 0, 0.5, 0),
		eda.LinearTrack(chipPads[20], rightPads[7], 0, 0.5, 0),
		eda.LinearTrack(chipPads[19], rightPads[8], 0, 0.5, 0),
		eda.LinearTrack(chipPads[18], rightPads[9], 0, 0.5, 0),

		eda.LinearTrack(chipPads[17], rightPads[10], 0, 1.2, 0),
		eda.LinearTrack(chipPads[16], rightPads[11], 0, 0, 0, 0.7, -0.5, -1.5, 0),
		eda.LinearTrack(chipPads[14], rightPads[12], 0, 0, 0, 0.7, -1, -3.8, 0),
	}

	Board_nRF52840 = &eda.Component{
		//ClearWidth: 0.25,

		Cuts: path.Paths{
			path.RoundRect(30, 18, 2),
		},

		Nested: eda.Components{
			boards.MountHole.CloneX(2, 24).Arrange(transform.Move(0, -6.8)),

			{
				Nested: eda.Components{
					boards.Logo.Arrange(transform.ScaleUniformly(2).Move(-9, -5.6)),
					eda.CenteredText("nRF52840").Arrange(transform.Scale(3.8, 1.6).Move(0, -8)),
					boards.Firm.Arrange(transform.ScaleUniformly(0.9).Move(9, -6)),
					boards.Rev(25, 10, 7).Arrange(transform.ScaleUniformly(0.9).Move(12, -8.4)),

					e73,
					leftConn,
					rightConn,

					{
						Tracks: append(
							leftTracks,
							rightTracks...,
						),
					},

					{
						ClearWidth:  eda.ClearOff,
						TracksWidth: 0.5,

						Tracks: path.Paths{
							eda.LinearTrack(chipPads[4], chipPads[4].Move(2, 0)),
							eda.LinearTrack(chipPads[15], chipPads[15].Move(0, 2)),
						},
					},
				},
			},
		},
	}
)
