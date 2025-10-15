// Copyright © 2025 Alex Temnok. All rights reserved.

package x1

import (
	"github.com/stretchr/testify/assert"
	"github.com/temnok/pcbc/eda"
	"github.com/temnok/pcbc/eda/boards"
	"github.com/temnok/pcbc/eda/boards/ebyte/e73"
	"github.com/temnok/pcbc/eda/pcb"
	"github.com/temnok/pcbc/eda/pcb/config"
	"github.com/temnok/pcbc/path"
	"testing"
)

func TestBoard(t *testing.T) {
	conf := config.Default()
	conf.Width = 36
	conf.Height = 20

	board := &eda.Component{
		Nested: eda.Components{
			{
				Cuts: path.Paths{
					path.RoundRect(35, 19, 1),
				},
			},

			boards.AlignHole.CloneX(2, 32.3).CloneY(2, 16.3),

			e73.Board_nRF52840,
		},
	}

	assert.NoError(t, pcb.Process(conf, board))
}
