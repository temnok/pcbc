// Copyright © 2025 Alex Temnok. All rights reserved.

package qfn16

import (
	"github.com/stretchr/testify/assert"
	"temnok/pcbc/eda/pcb"
	"testing"
)

func TestBoard(t *testing.T) {
	board := pcb.New(Board)
	board.SaveEtchOverride = board.SaveEtchPI
	assert.NoError(t, board.Process().SaveFiles())
}
