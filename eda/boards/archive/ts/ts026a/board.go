// Copyright © 2025 Alex Temnok. All rights reserved.

package ts026a

import (
	"temnok/pcbc/eda"
	"temnok/pcbc/eda/boards"
	"temnok/pcbc/eda/lib/header/mph100imp40f"
	"temnok/pcbc/eda/lib/ts/hyp"
	"temnok/pcbc/path"
	"temnok/pcbc/transform"
)

var Board = &eda.Component{
	Cuts: path.Paths{
		path.RoundRect(9, 5.5, 1),
	},

	Nested: eda.Components{
		hyp.Switch1TS026A.Arrange(transform.RotateDegrees(90).Move(3, 0)),
		mph100imp40f.G_V_SP(2).Arrange(transform.RotateDegrees(-90).Move(-3, 0)),
		boards.MountHole,

		boards.Logo.Arrange(transform.ScaleUniformly(0.7).Move(0, -1.9)),
		eda.CenteredText("SW").Arrange(transform.ScaleUniformly(1.5).Move(0, 2)),
	},
}

func init() {
	pad := Board.PadCenters()

	Board.Tracks = eda.Tracks(
		eda.Track{pad[0]}.XY(pad[1]),
		eda.Track{pad[1]}.DY(0.7).XY(pad[4]),

		eda.Track{pad[2]}.XY(pad[3]),
		eda.Track{pad[3]}.DY(-0.7).XY(pad[5]),
	)
}
