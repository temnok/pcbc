// Copyright © 2025 Alex Temnok. All rights reserved.

package qfn16

import (
	"github.com/stretchr/testify/assert"
	"temnok/pcbc/eda/pcb"
	"testing"
)

func TestBoard(t *testing.T) {
	Board.TrackWidth = 0.3
	board := pcb.New(Board)
	board.ExtraCopperWidth = 0
	board.CopperClearWidth = 0.2
	assert.NoError(t, board.Process().SaveFiles())
}
