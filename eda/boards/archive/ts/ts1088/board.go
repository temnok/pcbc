// Copyright © 2025 Alex Temnok. All rights reserved.

package ts1088

import (
	"github.com/temnok/pcbc/eda"
	"github.com/temnok/pcbc/eda/boards"
	"github.com/temnok/pcbc/eda/lib/header/mph100imp40f"
	"github.com/temnok/pcbc/eda/lib/ts/xunpu"
	"github.com/temnok/pcbc/path"
	"github.com/temnok/pcbc/transform"
)

var Board = &eda.Component{
	Cuts: path.Paths{
		path.RoundRect(9.5, 6.5, 1),
	},

	Nested: eda.Components{
		xunpu.SwitchTS1088.Arrange(transform.RotateDegrees(-90).Move(3, 0)),
		mph100imp40f.G_V_SP(2).Arrange(transform.RotateDegrees(-90).Move(-3.25, 0)),
		boards.MountHole.Arrange(transform.Move(-0.25, 0)),

		boards.Logo.Arrange(transform.Move(-1, -2.1)),
		eda.CenteredText("SW").Arrange(transform.Scale(2, 1.5).Move(-0.25, 2.4)),
	},
}
