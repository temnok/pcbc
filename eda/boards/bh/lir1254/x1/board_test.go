// Copyright © 2025 Alex Temnok. All rights reserved.

package x1

import (
	"github.com/temnok/pcbc/eda"
	"github.com/temnok/pcbc/eda/boards"
	"github.com/temnok/pcbc/eda/boards/bh/lir1254"
	"github.com/temnok/pcbc/eda/pcb"
	"github.com/temnok/pcbc/eda/pcb/config"
	"github.com/temnok/pcbc/path"
	"testing"
)

func TestBoard(t *testing.T) {
	conf := config.Default()
	conf.Width, conf.Height = 22, 20

	pcb.Process(conf, &eda.Component{
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
