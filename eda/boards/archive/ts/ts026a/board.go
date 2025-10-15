// Copyright © 2025 Alex Temnok. All rights reserved.

package ts026a

import (
	"github.com/temnok/pcbc/eda"
	"github.com/temnok/pcbc/eda/boards"
	"github.com/temnok/pcbc/eda/lib/header/mph100imp40f"
	"github.com/temnok/pcbc/eda/lib/ts/hyp"
	"github.com/temnok/pcbc/path"
	"github.com/temnok/pcbc/transform"
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
