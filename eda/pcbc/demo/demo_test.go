package demo

import (
	"github.com/stretchr/testify/assert"
	"temnok/pcbc/eda"
	"temnok/pcbc/eda/lib"
	"temnok/pcbc/eda/pcbc"
	"temnok/pcbc/eda/pcbc/py32f002a"
	"temnok/pcbc/geom"
	"testing"
)

func Test_Demo(t *testing.T) {
	for _, mode := range []string{"normal", "groundfill"} {
		pcb := eda.NewPCB(36, 46)

		pcb.Groundfill = mode == "groundfill"

		pcb.Component(&lib.Component{
			Components: lib.Components{
				pcbc.Board35x45,
				{
					Transform: geom.MoveXY(0, 9),
					Components: lib.Components{
						py32f002a.QFN16,
					},
				},
				{
					Transform: geom.MoveXY(-6.5, -7),
					Components: lib.Components{
						py32f002a.SOP8,
					},
				},
				{
					Transform: geom.MoveXY(6.5, -7),
					Components: lib.Components{
						py32f002a.SOP8,
					},
				},
				//{
				//	Transform: geom.MoveXY(6.5, -7).RotateD(90),
				//	Components: lib.Components{
				//		usbc.Board,
				//	},
				//},
			},
		})

		assert.NoError(t, pcb.SaveFiles("out/"+mode+"/"))
	}
}
