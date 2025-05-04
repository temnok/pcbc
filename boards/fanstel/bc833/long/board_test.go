// Copyright © 2025 Alex Temnok. All rights reserved.

package long

import (
	"github.com/stretchr/testify/assert"
	"temnok/pcbc/boards"
	"temnok/pcbc/boards/fanstel/bc833"
	"temnok/pcbc/eda"
	"temnok/pcbc/eda/pcb"
	"testing"
)

func Test_BC833(t *testing.T) {
	assert.NoError(t, pcb.ProcessWithDefaultConfig(&eda.Component{
		Components: eda.Components{
			boards.Board35x45,
			bc833.Board,
		},
	}))
}
