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
	"testing"
)

func TestBoard(t *testing.T) {
	conf := config.Default()
	conf.Width = 29
	conf.Height = 19

	assert.NoError(t, pcb.Process(conf, &eda.Component{
		Nested: eda.Components{
			{
				Cuts: path.Paths{
					path.RoundRect(28, 18, 1.5),
				},
			},

			boards.AlignHole.CloneX(2, 25).CloneY(2, 15),

			bc833.Board,
		},
	}))
}
