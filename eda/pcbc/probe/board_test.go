// Copyright © 2025 Alex Temnok. All rights reserved.

package qfn16

import (
	"github.com/stretchr/testify/assert"
	"temnok/pcbc/eda"
	"testing"
)

func TestBoard(t *testing.T) {
	board := &eda.Component{
		Tracks: eda.Tracks(
			eda.Track{{-5, 0.75}}.DX(10),
			eda.Track{{-5, 0.5}}.DX(10),
			eda.Track{{-5, 0.25}}.DX(10),
			eda.Track{{-5, 0}}.DX(10),
			eda.Track{{-5, -0.25}}.DX(10),
			eda.Track{{-5, -0.5}}.DX(10),
			eda.Track{{-5, -0.75}}.DX(10),
		),
	}

	pcb := eda.NewPCB(board)
	pcb.TrackWidth = 0.1
	pcb.ClearBrushDiameter = 0.25
	pcb.Process()
	assert.NoError(t, pcb.SaveFiles())
}
