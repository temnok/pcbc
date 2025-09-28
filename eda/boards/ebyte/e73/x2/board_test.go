// Copyright © 2025 Alex Temnok. All rights reserved.

package x2

import (
	"github.com/stretchr/testify/assert"
	"temnok/pcbc/eda"
	"temnok/pcbc/eda/boards"
	"temnok/pcbc/eda/boards/ebyte/e73"
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

	board := &eda.Component{
		Nested: eda.Components{
			{
				CutsOuter: true,

				Cuts: path.Paths{
					path.RoundRect(34, 44, 2),
				},
			},

			boards.AlignHole.CloneX(2, 30).CloneY(3, 20),

			e73.Board_nRF52840.Arrange(transform.Move(0, 10)),

			e73.Board_nRF52840.Arrange(transform.RotateDegrees(180).Move(0, -10)),
		},
	}

	assert.NoError(t, pcb.Process(conf, board))
}
