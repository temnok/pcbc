// Copyright © 2025 Alex Temnok. All rights reserved.

package e73

import (
	"github.com/stretchr/testify/assert"
	"temnok/pcbc/eda"
	"temnok/pcbc/eda/pcb"
	"temnok/pcbc/eda/pcbc"
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
