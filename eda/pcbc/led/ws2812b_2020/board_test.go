// Copyright © 2025 Alex Temnok. All rights reserved.

package ws2812b_2020

import (
	"github.com/stretchr/testify/assert"
	"temnok/pcbc/eda"
	"temnok/pcbc/eda/pcbc"
	"temnok/pcbc/eda/pcbc/fanstel/bc833"
	"temnok/pcbc/transform"
	"testing"
)

var testBoard = &eda.Component{
	Components: eda.Components{
		pcbc.Board35x45,
		bc833.Board.Arrange(transform.Move(0, 3.5)),
		eda.ComponentGrid(2, 11, 7,
			Board,
			Board,
		).Arrange(transform.Move(0, -13)),
	},
}

func TestBoard(t *testing.T) {
	assert.NoError(t, eda.GeneratePCB(testBoard))
}
