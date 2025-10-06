// Copyright © 2025 Alex Temnok. All rights reserved.

package p2

import (
	"temnok/pcbc/eda"
	"temnok/pcbc/eda/boards"
	"temnok/pcbc/eda/lib/pkg/smd"
	"temnok/pcbc/eda/pcb"
	"temnok/pcbc/eda/pcb/config"
	"temnok/pcbc/path"
	"temnok/pcbc/transform"
	"testing"
)

func TestBoard(t *testing.T) {
	conf := config.Default()
	conf.Width, conf.Height = 15, 19

	pcb.Process(conf, &eda.Component{
		Nested: eda.Components{
			{
				Cuts: path.Paths{
					path.RoundRect(14, 18, 1.4),
				},
			},

			boards.AlignHole.CloneX(2, 11).CloneY(2, 15),

			P2("R ", "1K0", smd.I0402).Arrange(transform.Move(0, 4.5)),
			P2("R ", "10K", smd.I0402).Arrange(transform.Move(0, 0)),
			P2("R ", "M10", smd.I0402).Arrange(transform.Move(0, -4.5)),
		},
	})
}
