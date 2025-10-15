// Copyright © 2025 Alex Temnok. All rights reserved.

package resistors

import (
	"github.com/temnok/pcbc/eda"
	"github.com/temnok/pcbc/eda/boards"
	"github.com/temnok/pcbc/eda/boards/p2"
	"github.com/temnok/pcbc/eda/boards/ts/ts026a"
	"github.com/temnok/pcbc/path"
)

var Board = &eda.Component{
	Nested: eda.Components{
		{
			Cuts: path.Paths{
				path.RoundRect(34, 40, 1.4),
			},
		},

		boards.AlignHole.CloneX(2, 30.5).CloneY(2, 36.5),

		eda.ComponentGrid(3, 10.9, 4.9,
			p2.P2_I0402("+LR", "-2V"),
			p2.P2_I0402("+LG", "-3V"),
			p2.P2_I0402("+LB", "-3V"),
			p2.P2_I0402("+LY", "-2V"),
			p2.P2_I0402("+LW", "-3V"),

			p2.P2_I0402(" R ", "51R"),
			p2.P2_I0402(" R ", "51R"),
			p2.P2_I0402(" R ", "75R"),
			p2.P2_I0402(" R ", "75R"),
			p2.P2_I0402(" R ", "K10"),
			p2.P2_I0402(" R ", "K10"),
			p2.P2_I0402(" R ", "K10"),
			p2.P2_I0402(" R ", "K15"),
			p2.P2_I0402(" R ", "K15"),
			p2.P2_I0402(" R ", "K20"),
			p2.P2_I0402(" R ", "K20"),

			ts026a.Board,
			ts026a.Board,
			ts026a.Board,
			ts026a.Board,
			ts026a.Board,
		),
	},
}
