// Copyright © 2025 Alex Temnok. All rights reserved.

package ws2812b_2020

import (
	"github.com/stretchr/testify/assert"
	"temnok/pcbc/eda"
	"temnok/pcbc/eda/boards"
	"temnok/pcbc/eda/pcb"
	"temnok/pcbc/eda/pcb/config"
	"temnok/pcbc/transform"
	"testing"
)

var testBoard = &eda.Component{
	Nested: eda.Components{
		boards.Board35x45,
		eda.ComponentGrid(2, 11, 7,
			Board,
			Board,
		).Arrange(transform.Move(0, -13)),
	},
}

func TestBoard(t *testing.T) {
	conf := config.Default()
	conf.Width, conf.Height = 36, 46

	assert.NoError(t, pcb.Process(conf, testBoard))
}
