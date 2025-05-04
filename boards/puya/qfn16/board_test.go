// Copyright © 2025 Alex Temnok. All rights reserved.

package qfn16

import (
	"github.com/stretchr/testify/assert"
	"temnok/pcbc/eda/pcb"
	"testing"
)

func TestBoard(t *testing.T) {
	Board.TrackWidth = 0.3
	config := pcb.Defaults()
	config.ExtraCopperWidth = 0
	config.CopperClearWidth = 0.2
	assert.NoError(t, pcb.SaveFiles(config, Board))
}
