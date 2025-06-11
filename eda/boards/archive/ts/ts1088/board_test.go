// Copyright © 2025 Alex Temnok. All rights reserved.

package ts1088

import (
	"github.com/stretchr/testify/assert"
	"temnok/pcbc/eda"
	"temnok/pcbc/eda/boards"
	"temnok/pcbc/eda/boards/archive/bh/lir1254"
	"temnok/pcbc/eda/pcb"
	"temnok/pcbc/transform"
	"testing"
)

var testBoard = &eda.Component{
	Inner: eda.Components{
		boards.Board35x45,
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
	assert.NoError(t, pcb.Process(nil, testBoard))
}
