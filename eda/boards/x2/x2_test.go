package x2

import (
	"temnok/pcbc/eda"
	"temnok/pcbc/eda/lib/pkg/smd"
	"temnok/pcbc/eda/pcb"
	"temnok/pcbc/path"
	"temnok/pcbc/transform"
	"testing"
)

func TestBoard(t *testing.T) {
	pcb.ProcessWithDefaultConfig(&eda.Component{
		Components: eda.Components{
			X2("R ", "1R0", smd.I0201),

			{
				OuterCut: true,
				Cuts: path.Paths{
					path.RoundRect(12, 7, 1),
					path.Circle(1).Transform(transform.Move(-5, 2.5)),
					path.Circle(1).Transform(transform.Move(5, -2.5)),
				},
			},
		},
	})
}
