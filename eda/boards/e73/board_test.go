// Copyright © 2025 Alex Temnok. All rights reserved.

package e73

import (
	"github.com/stretchr/testify/assert"
	"temnok/pcbc/eda"
	"temnok/pcbc/eda/pcb"
	"temnok/pcbc/eda/pcb/config"
	"temnok/pcbc/path"
	"temnok/pcbc/transform"
	"testing"
)

func TestBoard(t *testing.T) {
	conf := config.Default()
	conf.Width = 35
	conf.Height = 45
	conf.SavePath = "out/{}-"

	board := &eda.Component{
		Nested: eda.Components{
			{
				CutsOuter: true,

				Cuts: path.Paths{
					path.RoundRect(34, 44, 2),
				},
			},

			(&eda.Component{
				CutsOuter: true,
				CutsInner: true,

				Cuts: path.Paths{
					path.Circle(1.45),
				},
			}).Clone(2, 30, 0).Clone(3, 0, 20),

			Board_nRF52840.Arrange(transform.Move(0, 10)),

			Board_nRF52840.Arrange(transform.RotateDegrees(180).Move(0, -10)),
		},
	}

	assert.NoError(t, pcb.Process(conf,
		board.InLayer(1),
		board.InLayer(2).Arrange(transform.MirrorX()),
	))
}
