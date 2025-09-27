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
	e73      = ebyte.E73.Arrange(transform.Move(0, 3.07))
	chipPads = e73.PadCenters()

	leftConn = greenconn.CSCC118(14, false, []string{
		"P111", "P110", "P003", " AI4", "GND ", "P113", "AI0 ",
		" AI5", "AI7 ", " AI6", "XL1 ", " XL2", "AI3 ", "P109",
	}).Arrange(transform.Move(-11.25, 1.6))
	leftPads = leftConn.PadCenters()

	rightConn = greenconn.CSCC118(14, true, []string{
		" NF2", "NF1 ", " SWC", "SWD ", "P024", "P013", " D+ ",
		" D- ", "VBUS", "RST ", "DCCH", "VDDH", " GND", "VDD ",
	}).Arrange(transform.Move(11.25, 1.6))
	rightPads = rightConn.PadCenters()

	leftTracks = path.Paths{
		eda.Track(chipPads[0], leftPads[0]),
		eda.Track(chipPads[1], leftPads[1], 0, -0.5, 0),
		eda.Track(chipPads[2], leftPads[2], 0, -0.5, 0),
		eda.Track(chipPads[3], leftPads[3], 0, -0.5, 0),
		eda.Track(chipPads[4], leftPads[4], 0, -0.5, 0),
		eda.Track(chipPads[5], leftPads[5], 0, -0.5, 0),
		eda.Track(chipPads[6], leftPads[6], 0, -0.5, 0),
		eda.Track(chipPads[7], leftPads[7], 0, -0.5, 0),
		eda.Track(chipPads[8], leftPads[8], 0, -0.5, 0),
		eda.Track(chipPads[9], leftPads[9], 0, -0.5, 0),

		eda.Track(chipPads[10], leftPads[10], 0, -1.2, 0),
		eda.Track(chipPads[11], leftPads[11], 0, -1e-9, -1e-9, -0.7, 0.5, 1.5, 0),
		eda.Track(chipPads[12], leftPads[12], 0, -1e-9, -1e-9, -0.7, 1.0, 2.5, 0),
		eda.Track(chipPads[13], leftPads[13], 0, -1e-9, -1e-9, -0.7, 1.5, 3.5, 0),
	}

	Board_nRF52840 = &eda.Component{
		ClearWidth: 0.25,

		Cuts: path.Paths{
			path.RoundRect(29, 19, 1.5),
		},

		Nested: eda.Components{
			boards.MountHole.Clone(2, 20, 0).Arrange(transform.Move(0, -7.2)),

			{
				Nested: eda.Components{
					boards.Logo.Arrange(transform.Scale(1.6, 1.6).Move(-13, -7.1)),
					eda.CenteredText("E73-2G4M08S1C").Arrange(transform.Scale(2, 1.6).Move(0, -7.5)),
					eda.CenteredText("nRF52840").Arrange(transform.Scale(2, 1.2).Move(0, -8.7)),
					boards.Firm.Arrange(transform.Scale(0.9, 0.9).Move(13, -7.2)),

					e73,
					leftConn,
					rightConn,

					{
						Tracks: append(
							leftTracks,
							leftTracks.Transform(transform.MirrorX())...,
						),
					},

					{
						ClearOff:    true,
						TracksWidth: 0.3,

						Tracks: path.Paths{
							eda.Track(chipPads[4], chipPads[4].Move(2, 0)),
							eda.Track(leftPads[4], leftPads[4].Move(-2, 0)),

							eda.Track(chipPads[15], chipPads[15].Move(0, 2)),
							eda.Track(rightPads[12], rightPads[12].Move(2, 0)),
						},
					},
				},
			},
		},
	}
)
