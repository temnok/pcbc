package bc833

import (
	"github.com/stretchr/testify/assert"
	"temnok/pcbc/eda"
	"temnok/pcbc/eda/lib"
	"temnok/pcbc/eda/pcbc"
	"temnok/pcbc/geom"
	"temnok/pcbc/path"
	"testing"
)

func Test_BC833Normal(t *testing.T) {
	pcb := eda.NewPCB(36, 46)
	pcb.Component(&lib.Component{
		Clears: path.Paths{path.Rect(36, 46)},
		Components: lib.Components{
			pcbc.Board35x45,
			Board,
		},
	})
	assert.NoError(t, pcb.SaveFiles("out/normal/"))
}

func Test_BC833NormalShort(t *testing.T) {
	pcb := eda.NewPCB(36, 46)
	pcb.Component(shortBoard(true))
	assert.NoError(t, pcb.SaveFiles("out/normal-short/"))
}

func Test_BC833(t *testing.T) {
	pcb := eda.NewPCB(36, 46)
	pcb.Component(&lib.Component{
		Components: lib.Components{
			pcbc.Board35x45,
			Board,
		},
	})
	assert.NoError(t, pcb.SaveFiles("out/ground/"))
}

func Test_BC833Short(t *testing.T) {
	pcb := eda.NewPCB(36, 46)
	pcb.Component(shortBoard(false))
	assert.NoError(t, pcb.SaveFiles("out/ground-short/"))
}

func shortBoard(clear bool) *lib.Component {
	clears := path.Paths{path.Rect(36, 46)}
	if !clear {
		clears = nil
	}

	return &lib.Component{
		Clears: clears,
		Components: lib.Components{
			pcbc.Board35x45,
			{
				Transform: geom.MoveXY(0, 10.5),
				Components: lib.Components{
					ShortBoard,
				},
			},
			{
				Transform: geom.MoveXY(0, -10.5).RotateD(180),
				Components: lib.Components{
					ShortBoard,
				},
			},
		},
	}
}
