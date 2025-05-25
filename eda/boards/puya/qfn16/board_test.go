// Copyright © 2025 Alex Temnok. All rights reserved.

package qfn16

import (
	"github.com/stretchr/testify/assert"
	"temnok/pcbc/eda/pcb"
	"temnok/pcbc/eda/pcb/config"
	"testing"
)

func TestBoard(t *testing.T) {
	Board.TrackWidth = 0.15
	config := config.Default()
	config.ExtraCopperWidth = 0.0
	//config.ExtraCopperWidth = 0.025
	config.ExtraPadCopperWidth = 0.15
	config.CopperClearWidth = 0.1

	//config.SpacerPeriod = 0.30
	//config.SpacerWidth = 0.15

	assert.NoError(t, pcb.Process(config, Board))
}
