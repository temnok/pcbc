package ts1088

import (
	"github.com/stretchr/testify/assert"
	"temnok/pcbc/eda"
	"temnok/pcbc/eda/lib"
	"temnok/pcbc/eda/pcbc"
	"temnok/pcbc/eda/pcbc/bh/lir1254"
	"temnok/pcbc/transform"
	"testing"
)

var testBoard = &lib.Component{
	Components: lib.Components{
		pcbc.Board35x45,
		lir1254.Board.Arrange(transform.Move(0, 8)),
		lib.ComponentGrid(2, 11, 8,
			Board,
			Board,
			Board,
			Board,
		).Arrange(transform.Move(0, -8)),
	},
}

func TestBoard(t *testing.T) {
	assert.NoError(t, eda.GeneratePCB(testBoard))
}
