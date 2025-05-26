// Copyright © 2025 Alex Temnok. All rights reserved.

package qfn16

import (
	"github.com/stretchr/testify/assert"
	"temnok/pcbc/eda/pcb"
	"temnok/pcbc/eda/pcb/config"
	"testing"
)

func TestBoard(t *testing.T) {
	Board.TrackWidth = 0.25
	config := config.Default()
	config.ExtraCopperWidth = 0.0
	config.CopperClearWidth = 0.15

	assert.NoError(t, pcb.Process(config, Board))
}
