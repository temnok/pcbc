// Copyright © 2025 Alex Temnok. All rights reserved.

package lir1254

import (
	"github.com/stretchr/testify/assert"
	"temnok/pcbc/eda"
	"temnok/pcbc/eda/pcb"
	"temnok/pcbc/eda/pcb/config"
	"temnok/pcbc/path"
	"temnok/pcbc/transform"
	"testing"
)

var hole = path.Circle(1.45)

var blank = &eda.Component{
	CutsOuter: true,

	Cuts: path.Paths{
		path.RoundRect(24, 24, 2),

		hole.Transform(transform.Move(-10, -10)),
		hole.Transform(transform.Move(-10, 10)),
		hole.Transform(transform.Move(10, -10)),
		hole.Transform(transform.Move(10, 10)),
	},
}

var testBoard = &eda.Component{
	Inner: eda.Components{
		blank,
		Board.Arrange(transform.Move(0, 1)),
	},
}

func TestBoard(t *testing.T) {
	conf := config.Default()
	conf.Width, conf.Height = 25, 25

	assert.NoError(t, pcb.Process(conf, testBoard))
}
