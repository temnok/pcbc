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
	"testing"
)

func TestBoard(t *testing.T) {
	conf := config.Default()
	conf.Width = 29
	conf.Height = 19

	assert.NoError(t, pcb.Process(conf, &eda.Component{
		Nested: eda.Components{
			{
				CutsOuter: true,

				Cuts: path.Paths{
					path.RoundRect(28, 18, 1.5),
				},
			},

			boards.AlignHole.Clone(2, 25, 0).Clone(2, 0, 15),

			bc833.Board,
		},
	}))
}

/*

3x2: 22x20
2x2: 35x20
3x1: 22x45
2x1: 36x45
1x1: 77x45

*/
