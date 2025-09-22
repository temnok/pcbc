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
	e73     = ebyte.E73.Arrange(transform.Move(0, 3.07))
	e73pads = e73.PadCenters()

	leftConn = greenconn.CSCC118(14, false, []string{
		"P111", "P110", "P003", " AI4", "GND ", "P113", "AI0 ", " AI5", "AI7 ", " AI6",
		"XL1 ", " XL2", "AI3 ", "P109",
	}).Arrange(transform.Move(-11.25, 1.6))
	leftConnPads = leftConn.PadCenters()

	rightConn = greenconn.CSCC118(14, true, []string{
		" NF2", "NF1 ", " SWC", "SWD ", "P024", "P013", "  D+", "D-  ", " VBS", "RST ",
		" DCH", "VDH ", " GND", "VDD ",
	}).Arrange(transform.Move(11.25, 1.6))
	rightConnPads = rightConn.PadCenters()

	leftTracks = eda.Tracks(
		eda.TrackV2(e73pads[0], -0.5, 0, 0.00, -3.5),
		eda.TrackV2(e73pads[1], -0.5, 0, 0.27, -5.5),
		eda.TrackV2(e73pads[2], -0.5, 0, 0.54, -3.0),
		eda.TrackV2(e73pads[3], -0.5, 0, 0.81, -5.0),
		eda.TrackV2(e73pads[4], -0.5, 0, 1.08, -2.5),
		eda.TrackV2(e73pads[5], -0.5, 0, 1.35, -4.5),
		eda.TrackV2(e73pads[6], -0.5, 0, 1.62, -2.0),
		eda.TrackV2(e73pads[7], -0.5, 0, 1.89, -4.0),
		eda.TrackV2(e73pads[8], -0.5, 0, 2.16, -1.5),
		eda.TrackV2(e73pads[9], -0.5, 0, 2.43, -3.5),

		eda.TrackV2(e73pads[10], -1.2, 0, 4, -1),
		eda.TrackV2(e73pads[11], 0, -0.7, 0.5, 1.5, 4.25, -3),
		eda.TrackV2(e73pads[12], 0, -0.7, 1.0, 2.5, 3.75, -1),
		eda.TrackV2(e73pads[13], 0, -0.7, 1.5, 3.5, 3.25, -3),
	)

	Board_nRF52840 = &eda.Component{
		ClearWidth: 0.25,

		Cuts: path.Paths{
			path.RoundRect(29, 19, 1.5),
		},

		Nested: eda.Components{
			boards.MountHole15(1).Arrange(transform.Move(-10, -7.2)),
			boards.MountHole15(1).Arrange(transform.Move(10, -7.2)),

			{
				Layer: 1,

				Nested: eda.Components{
					boards.Logo.Arrange(transform.Scale(1.6, 1.6).Move(-13, -7.1)),
					eda.CenteredText("E73-2G4M08S1C").Arrange(transform.Scale(2, 1.6).Move(0, -7.4)),
					eda.CenteredText("nRF52840").Arrange(transform.Scale(2, 1.2).Move(0, -8.6)),
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
						ClearDisabled: true,
						TracksWidth:   0.3,

						Tracks: eda.Tracks(
							eda.TrackV2(e73pads[4], 2, 0),
							eda.TrackV2(leftConnPads[4], -2, 0),

							eda.TrackV2(e73pads[15], 0, 2),
							eda.TrackV2(rightConnPads[12], 2, 0),
						),
					},
				},
			},
		},
	}
)
