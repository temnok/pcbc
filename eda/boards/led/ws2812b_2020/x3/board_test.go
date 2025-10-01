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
	conf.Width, conf.Height = 16, 20

	pcb.Process(conf, &eda.Component{
		Nested: eda.Components{
			{
				Cuts: path.Paths{
					path.RoundRect(15, 19, 1.4),
				},
			},

			boards.AlignHole.CloneX(2, 12).CloneY(2, 16),

			ws2812b_2020.Board.CloneY(3, 5),
		},
	})
}
