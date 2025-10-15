// Copyright © 2025 Alex Temnok. All rights reserved.

package ts1088

import (
	"github.com/stretchr/testify/assert"
	"github.com/temnok/pcbc/eda"
	"github.com/temnok/pcbc/eda/boards/bh/lir1254"
	"github.com/temnok/pcbc/eda/pcb"
	"github.com/temnok/pcbc/eda/pcb/config"
	"github.com/temnok/pcbc/transform"
	"testing"
)

var testBoard = &eda.Component{
	Nested: eda.Components{
		lir1254.Board.Arrange(transform.Move(0, 8)),
		eda.ComponentGrid(2, 11, 8,
			Board,
			Board,
			Board,
			Board,
		).Arrange(transform.Move(0, -8)),
	},
}

func TestBoard(t *testing.T) {
	conf := config.Default()
	conf.Width, conf.Height = 36, 46

	assert.NoError(t, pcb.Process(conf, testBoard))
}
