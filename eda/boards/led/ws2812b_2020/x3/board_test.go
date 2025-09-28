// Copyright © 2025 Alex Temnok. All rights reserved.

package x3

import (
	"temnok/pcbc/eda"
	"temnok/pcbc/eda/boards"
	"temnok/pcbc/eda/boards/led/ws2812b_2020"
	"temnok/pcbc/eda/pcb"
	"temnok/pcbc/eda/pcb/config"
	"temnok/pcbc/path"
	"testing"
)

func TestBoard(t *testing.T) {
	conf := config.Default()
	conf.Width, conf.Height = 14, 19

	pcb.Process(conf, &eda.Component{
		Nested: eda.Components{
			{
				CutsOuter: true,

				Cuts: path.Paths{
					path.RoundRect(13, 18, 1.4),
				},
			},

			boards.AlignHole.CloneX(2, 10).CloneY(2, 15),

			ws2812b_2020.Board.CloneY(3, 4.4),
		},
	})
}
