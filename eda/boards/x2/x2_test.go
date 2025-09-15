// Copyright © 2025 Alex Temnok. All rights reserved.

package x2

import (
	"temnok/pcbc/eda"
	"temnok/pcbc/eda/lib/pkg/smd"
	"temnok/pcbc/eda/pcb"
	"temnok/pcbc/eda/pcb/config"
	"temnok/pcbc/path"
	"temnok/pcbc/transform"
	"testing"
)

func TestBoard(t *testing.T) {
	conf := config.Default()
	conf.Width, conf.Height = 13, 13

	//conf.ExtraCopperWidth = 0
	//conf.StencilPadOffset = 0.05

	hole := path.Circle(1)

	pcb.Process(conf, &eda.Component{
		Inner: eda.Components{
			X2("R ", "10K", smd.I0201).Arrange(transform.Move(0, 2.25)),
			X2("R ", "1K0", smd.I0402).Arrange(transform.Move(0, -2.25)),

			{
				CutsOuter: true,
				Cuts: path.Paths{
					path.RoundRect(12, 12, 1),

					hole.Transform(transform.Move(-5, -5)),
					hole.Transform(transform.Move(-5, 5)),
					hole.Transform(transform.Move(5, -5)),
					hole.Transform(transform.Move(5, 5)),
				},
			},
		},
	})
}
