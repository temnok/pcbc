// Copyright © 2025 Alex Temnok. All rights reserved.

package ws2812b_2020

import (
	"temnok/pcbc/eda"
	"temnok/pcbc/eda/boards"
	"temnok/pcbc/eda/lib/header/mph100imp40f"
	"temnok/pcbc/eda/lib/worldsemi"
	"temnok/pcbc/path"
	"temnok/pcbc/transform"
)

var Board = &eda.Component{
	Cuts: path.Paths{
		path.RoundRect(10, 5.5, 1),
	},

	Nested: eda.Components{
		worldsemi.WS2812B_2020.Arrange(transform.Move(-3, 0)),

		mph100imp40f.G_V_SP(2).Arrange(transform.RotateDegrees(-90).Move(3.5, 0)),

		boards.MountHole.Arrange(transform.RotateDegrees(45).Move(0.55, 0)),

		boards.Logo.Arrange(transform.ScaleUniformly(0.7).Move(-0.6, -2)),

		eda.CenteredText("WS2812B").Arrange(transform.Scale(1, 1.8).Move(-2.1, 1.8)),

		eda.CenteredText("LED").Arrange(transform.Scale(1.5, 1.8).Move(-3, -1.8)),

		eda.CenteredText("DI").Arrange(transform.Scale(0.8, 1.4).Move(1.6, 2)),

		eda.CenteredText("VDD").Arrange(transform.Scale(0.8, 1.4).Move(1.3, -2)),
	},
}

func init() {
	pad := Board.PadCenters()

	Board.Tracks = eda.Tracks(
		eda.Track{pad[3]}.DY(1.5).XY(pad[5]),
		eda.Track{pad[4]}.DY(-1.5).XY(pad[6]),
	)

	Board.Nested = append(Board.Nested, &eda.Component{
		ClearOff: true,

		Tracks: eda.Tracks(
			eda.Track{pad[2]}.DY(0.5).XY(pad[9]),
		),
	})
}
