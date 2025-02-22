// Copyright © 2025 Alex Temnok. All rights reserved.

package ts1088

import (
	"temnok/pcbc/eda"
	"temnok/pcbc/eda/lib/header/mph100imp40f"
	"temnok/pcbc/eda/lib/ts/xunpu"
	"temnok/pcbc/eda/pcbc"
	"temnok/pcbc/path"
	"temnok/pcbc/transform"
)

var Board = &eda.Component{
	Cuts: path.Paths{
		path.RoundRect(9.5, 6.5, 1),
	},

	Components: eda.Components{
		xunpu.SwitchTS1088.Arrange(transform.Rotate(-90).Move(3, 0)),
		mph100imp40f.G_V_SP(2).Arrange(transform.Rotate(-90).Move(-3.25, 0)),
		pcbc.MountHole.Arrange(transform.Move(-0.25, 0)),

		pcbc.Logo.Arrange(transform.Move(-1, -2.1)),
		eda.CenteredText("SW").Arrange(transform.Scale(2, 1.5).Move(-0.25, 2.4)),
	},
}

func init() {
	pad := Board.PadCenters()

	Board.Tracks = eda.Tracks(
		eda.Track{pad[0]}.XY(pad[2]),
		eda.Track{pad[1]}.XY(pad[3]),
	)
}
