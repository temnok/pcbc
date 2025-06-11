// Copyright © 2025 Alex Temnok. All rights reserved.

package lir1254

import (
	"github.com/stretchr/testify/assert"
	"temnok/pcbc/eda"
	"temnok/pcbc/eda/boards"
	"temnok/pcbc/eda/pcb"
	"testing"
)

var testBoard = &eda.Component{
	Components: eda.Components{
		boards.Board35x45,
		Board,
	},
}

func TestBoard(t *testing.T) {
	assert.NoError(t, pcb.Process(nil, testBoard))
}
