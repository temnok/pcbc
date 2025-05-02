// Copyright © 2025 Alex Temnok. All rights reserved.

package e73tiny

import (
	"github.com/stretchr/testify/assert"
	"temnok/pcbc/eda"
	"temnok/pcbc/eda/board"
	"temnok/pcbc/eda/pcb"
	"testing"
)

func TestBoard(t *testing.T) {
	assert.NoError(t, pcb.Generate(&eda.Component{
		Components: eda.Components{
			pcbc.Board35x45,
			Board_nRF52840,
		},
	}))
}
