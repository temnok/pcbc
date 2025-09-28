// Copyright © 2025 Alex Temnok. All rights reserved.

package x2

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
	conf.Width, conf.Height = 14, 14

	pcb.Process(conf, &eda.Component{
		Nested: eda.Components{
			{
				CutsOuter: true,

				Cuts: path.Paths{
					path.RoundRect(13, 13, 1.4),
				},
			},

			boards.AlignHole.CloneX(2, 10).CloneY(2, 10),

			X2("R ", "10K", smd.I0402).Arrange(transform.Move(0, 2.1)),
			X2("R ", "1K0", smd.I0402).Arrange(transform.Move(0, -2.1)),
		},
	})
}
