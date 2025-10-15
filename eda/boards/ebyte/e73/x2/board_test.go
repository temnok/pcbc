// Copyright © 2025 Alex Temnok. All rights reserved.

package x2

import (
	"github.com/stretchr/testify/assert"
	"github.com/temnok/pcbc/eda"
	"github.com/temnok/pcbc/eda/boards"
	"github.com/temnok/pcbc/eda/boards/ebyte/e73"
	"github.com/temnok/pcbc/eda/pcb"
	"github.com/temnok/pcbc/eda/pcb/config"
	"github.com/temnok/pcbc/path"
	"github.com/temnok/pcbc/transform"
	"testing"
)

func TestBoard(t *testing.T) {
	conf := config.Default()
	conf.Width = 34
	conf.Height = 43

	board := &eda.Component{
		Nested: eda.Components{
			{
				Cuts: path.Paths{
					path.RoundRect(33, 42, 1.4),
				},
			},

			boards.AlignHole.CloneX(2, 30).CloneY(3, 19.5),

			e73.Board_nRF52840.Arrange(transform.Move(0, 9.75)),

			e73.Board_nRF52840.Arrange(transform.RotateDegrees(180).Move(0, -9.75)),
		},
	}

	assert.NoError(t, pcb.Process(conf, board))
}
