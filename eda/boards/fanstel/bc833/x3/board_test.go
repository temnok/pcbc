// Copyright © 2025 Alex Temnok. All rights reserved.

package x3

import (
	"github.com/stretchr/testify/assert"
	"github.com/temnok/pcbc/eda"
	"github.com/temnok/pcbc/eda/boards"
	"github.com/temnok/pcbc/eda/boards/fanstel/bc833"
	"github.com/temnok/pcbc/eda/pcb"
	"github.com/temnok/pcbc/eda/pcb/config"
	"github.com/temnok/pcbc/path"
	"github.com/temnok/pcbc/transform"
	"testing"
)

func TestBoard(t *testing.T) {
	conf := config.Default()
	conf.Width = 35
	conf.Height = 45

	board := &eda.Component{
		Nested: eda.Components{
			{
				Cuts: path.Paths{
					path.RoundRect(34, 44, 2),
				},
			},

			boards.AlignHole.CloneX(2, 30).CloneY(2, 40),

			bc833.Board.Arrange(transform.Move(0, 13)),

			bc833.Board.Arrange(transform.RotateDegrees(90).Move(-8, -7)),

			bc833.Board.Arrange(transform.RotateDegrees(-90).Move(8, -7)),
		},
	}

	assert.NoError(t, pcb.Process(conf, board))
}
