package demo

import (
	"github.com/stretchr/testify/assert"
	"temnok/pcbc/eda"
	"temnok/pcbc/eda/lib"
	"temnok/pcbc/eda/pcbc"
	"temnok/pcbc/eda/pcbc/py32f002a"
	"temnok/pcbc/geom"
	"temnok/pcbc/path"
	"testing"
)

var demoBoard = &lib.Component{
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
}

func Test_DemoGroundfill(t *testing.T) {
	pcb := eda.NewPCB(36, 46)
	pcb.Component(demoBoard)

	assert.NoError(t, pcb.SaveFiles("out/ground/"))
}

func Test_DemoNormal(t *testing.T) {
	pcb := eda.NewPCB(36, 46)
	pcb.Component(&lib.Component{
		Clears: path.Paths{
			path.Rect(36, 46),
		},

		Components: lib.Components{
			demoBoard,
		},
	})

	assert.NoError(t, pcb.SaveFiles("out/normal/"))
}
