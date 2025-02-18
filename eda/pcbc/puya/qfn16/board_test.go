// Copyright © 2025 Alex Temnok. All rights reserved.

package qfn16

import (
	"github.com/stretchr/testify/assert"
	"temnok/pcbc/eda"
	"testing"
)

func TestBoard(t *testing.T) {
	pcb := eda.NewPCB(Board)
	//pcb.TrackWidth = 0.35
	//pcb.ClearBrushDiameter = 0.2
	pcb.Process()
	assert.NoError(t, pcb.SaveFiles())
}
