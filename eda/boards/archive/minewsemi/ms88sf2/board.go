// Copyright © 2025 Alex Temnok. All rights reserved.

package ms88sf2

import (
	"temnok/pcbc/eda"
	"temnok/pcbc/eda/boards"
	"temnok/pcbc/eda/lib/header/mph100imp40f"
	"temnok/pcbc/eda/lib/minewsemi"
	"temnok/pcbc/font"
	"temnok/pcbc/path"
	"temnok/pcbc/transform"
)

var (
	labelShift = 2.54 / 0.8
	labelScale = transform.Scale(0.8, 1.8)

	chip = minewsemi.MS88SF2.Arrange(transform.Move(0, 6.6))

	pin = chip.PadCenters()

	headers = &eda.Component{
		Transform: transform.Move(0, 3.05),

		Nested: eda.Components{
			mph100imp40f.G_V_SP(8).Arrange(transform.RotateDegrees(-90).Move(-12.7, -1)),
			mph100imp40f.G_V_SP(11).Arrange(transform.Move(0, -14)),
			mph100imp40f.G_V_SP(8).Arrange(transform.RotateDegrees(90).Move(12.7, -1)),
		},

		Marks: path.Join(
			font.CenteredRow(labelShift,
				"P113", "P115", "P002", "P029", "P031", "P109", "P012", "GND",
			).Transform(labelScale.RotateDegrees(-90).Move(-10.6, -1)),
			font.CenteredRow(labelShift,
				"VDD", "P008", "P006", "P004", "P026", "P024", "P022", "P020", "P018", "P015", "VDDH",
			).Transform(labelScale.Move(0, -11.95)),
			font.CenteredRow(labelShift,
				"D-", "D+", "P013", "P100", "SWD", "SWC", "P009", "P010",
			).Transform(labelScale.RotateDegrees(90).Move(10.6, -1)),
		),
	}

	pad = append([]path.Point{{}}, headers.PadCenters()...)

	mountHoles = &eda.Component{
		Transform: transform.Move(0, 3),

		Nested: eda.Components{
			boards.MountHole.Arrange(transform.RotateDegrees(45).Move(-7.5, -9.7)),
			boards.MountHole.Arrange(transform.RotateDegrees(-45).Move(7.5, -9.7)),
		},
	}

	Board_nRF52840 = &eda.Component{
		Cuts: path.Paths{
			path.RoundRect(28.4, 24.9, 1),
		},

		Nested: eda.Components{
			chip,
			headers,
			mountHoles,

			boards.Logo.Arrange(transform.Move(-4.8, -6.5)),
			eda.CenteredText("MS88SF21").Arrange(transform.Scale(1.8, 1.8).Move(1.3, -5.9)),
			eda.CenteredText("nRF52840").Arrange(transform.ScaleUniformly(1.8).Move(1.3, -7.4)),

			{
				ClearOff: true,

				Tracks: eda.DeprecatedTracks(
					eda.DeprecatedTrack{pin[1]}.DX(1.5),
					eda.DeprecatedTrack{pin[1]}.DX(-1.5),
					eda.DeprecatedTrack{pad[8]}.DX(1.7).Y(pin[13].Y-1.4).YX(pin[13]),
				),
			},
		},

		Tracks: eda.DeprecatedTracks(
			eda.DeprecatedTrack{pad[1]}.YX(pin[2]),
			eda.DeprecatedTrack{pad[2]}.YX(pin[3]),
			eda.DeprecatedTrack{pad[3]}.YX(pin[4]),
			eda.DeprecatedTrack{pad[4]}.DX(1.7).Y(pin[5].Y-1.2).YX(pin[5]),
			eda.DeprecatedTrack{pad[5]}.DX(2.4).Y(pin[6].Y-0.4).YX(pin[6]),
			eda.DeprecatedTrack{pad[6]}.YX(pin[11]),
			eda.DeprecatedTrack{pad[7]}.YX(pin[12]),
			eda.DeprecatedTrack{pad[9]}.DX(2.4).Y(pin[14].Y-0.6).YX(pin[14]),

			eda.DeprecatedTrack{pad[10]}.DY(1.7).X(-4.8).DY(2.1).Y(pin[19].Y-1).YX(pin[10]),
			eda.DeprecatedTrack{pad[11]}.X(-3.6).Y(pin[20].Y-1.5).YX(pin[9]),
			eda.DeprecatedTrack{pad[12]}.X(-2.4).Y(pin[21].Y-2).YX(pin[8]),
			eda.DeprecatedTrack{pad[13]}.X(-1.2).Y(pin[22].Y-2.5).YX(pin[7]),

			eda.DeprecatedTrack{pad[14]}.Y(pin[23].Y-3).YX(pin[23]),

			eda.DeprecatedTrack{pad[15]}.X(1.2).Y(pin[22].Y-2.5).YX(pin[22]),
			eda.DeprecatedTrack{pad[16]}.X(2.4).Y(pin[21].Y-2).YX(pin[21]),
			eda.DeprecatedTrack{pad[17]}.X(3.6).Y(pin[20].Y-1.5).YX(pin[20]),
			eda.DeprecatedTrack{pad[18]}.DY(1.7).X(4.8).DY(2.1).Y(pin[19].Y-1).YX(pin[19]),

			eda.DeprecatedTrack{pad[19]}.DX(-2.9).Y(pin[15].Y-0.6).YX(pin[15]),
			eda.DeprecatedTrack{pad[20]}.DX(-1.8).Y(pin[16].Y-1.4).YX(pin[16]),
			eda.DeprecatedTrack{pad[21]}.YX(pin[17]),
			eda.DeprecatedTrack{pad[22]}.YX(pin[18]),
			eda.DeprecatedTrack{pad[23]}.DX(-2.5).Y(pin[24].Y-0.4).YX(pin[24]),
			eda.DeprecatedTrack{pad[24]}.DX(-1.8).Y(pin[25].Y-1.2).YX(pin[25]),
			eda.DeprecatedTrack{pad[25]}.YX(pin[26]),
			eda.DeprecatedTrack{pad[26]}.YX(pin[27]),
			eda.DeprecatedTrack{pad[27]}.YX(pin[28]),
		),
	}
)
