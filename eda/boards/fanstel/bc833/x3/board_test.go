// Copyright © 2025 Alex Temnok. All rights reserved.

package x3

import (
	"github.com/stretchr/testify/assert"
	"temnok/pcbc/eda"
	"temnok/pcbc/eda/boards"
	"temnok/pcbc/eda/boards/fanstel/bc833"
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

			boards.AlignHole.Clone(2, 30, 0).Clone(2, 0, 40),

			bc833.Board.Arrange(transform.Move(0, 13)),

			bc833.Board.Arrange(transform.RotateDegrees(90).Move(-8, -7)),

			bc833.Board.Arrange(transform.RotateDegrees(-90).Move(8, -7)),
		},
	}

	assert.NoError(t, pcb.Process(conf, board))
}
