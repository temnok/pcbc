// Copyright © 2025 Alex Temnok. All rights reserved.

package x1

import (
	"temnok/pcbc/eda"
	"temnok/pcbc/eda/boards"
	"temnok/pcbc/eda/boards/bh/lir1254"
	"temnok/pcbc/eda/pcb"
	"temnok/pcbc/eda/pcb/config"
	"temnok/pcbc/path"
	"testing"
)

func TestBoard(t *testing.T) {
	conf := config.Default()
	conf.Width, conf.Height = 22, 20

	pcb.Process(conf, &eda.Component{
		ClearWidth: 0.25,

		Nested: eda.Components{
			{
				Cuts: path.Paths{
					path.RoundRect(21, 19, 1.4),
				},
			},

			boards.AlignHole.CloneX(2, 18).CloneY(2, 16),

			lir1254.Board,
		},
	})
}
