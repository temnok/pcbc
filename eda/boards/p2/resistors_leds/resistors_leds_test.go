// Copyright © 2025 Alex Temnok. All rights reserved.

package resistors

import (
	"github.com/stretchr/testify/assert"
	"temnok/pcbc/eda"
	"temnok/pcbc/eda/boards"
	"temnok/pcbc/eda/boards/p2"
	"temnok/pcbc/eda/boards/ts/ts026a"
	"temnok/pcbc/eda/pcb"
	"temnok/pcbc/eda/pcb/config"
	"temnok/pcbc/path"
	"testing"
)

func TestBoard(t *testing.T) {
	conf := config.Default()
	conf.Width, conf.Height = 35, 41

	assert.NoError(t, pcb.Process(conf, &eda.Component{
		Nested: eda.Components{
			{
				Cuts: path.Paths{
					path.RoundRect(34, 40, 1.2),
				},
			},

			boards.AlignHole.CloneX(2, 31).CloneY(2, 37),

			eda.ComponentGrid(3, 11, 5,
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
	}))
}
