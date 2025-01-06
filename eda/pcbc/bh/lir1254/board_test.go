// Copyright © 2025 Alex Temnok. All rights reserved.

package lir1254

import (
	"github.com/stretchr/testify/assert"
	"temnok/pcbc/eda"
	"temnok/pcbc/eda/pcbc"
	"testing"
)

var testBoard = &eda.Component{
	//Clears: path.Paths{path.Rect(36, 46)},
	Components: eda.Components{
		pcbc.Board35x45,
		Board,
	},
}

func TestBoard(t *testing.T) {
	assert.NoError(t, eda.GeneratePCB(testBoard))
}
