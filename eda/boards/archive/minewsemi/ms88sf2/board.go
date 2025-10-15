// Copyright © 2025 Alex Temnok. All rights reserved.

package ms88sf2

import (
	"github.com/temnok/pcbc/eda"
	"github.com/temnok/pcbc/eda/boards"
	"github.com/temnok/pcbc/eda/lib/header/mph100imp40f"
	"github.com/temnok/pcbc/eda/lib/minewsemi"
	"github.com/temnok/pcbc/font"
	"github.com/temnok/pcbc/path"
	"github.com/temnok/pcbc/transform"
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
			boards.MountHole.CloneX(2, 16).Arrange(transform.Move(0, -9.7)),
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
				ClearWidth: eda.ClearOff,
			},
		},
	}
)
