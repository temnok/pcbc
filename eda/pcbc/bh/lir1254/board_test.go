package lir1254

import (
	"github.com/stretchr/testify/assert"
	"temnok/pcbc/eda"
	"temnok/pcbc/eda/lib"
	"temnok/pcbc/eda/pcbc"
	"testing"
)

var testBoard = &lib.Component{
	//Clears: path.Paths{path.Rect(36, 46)},
	Components: lib.Components{
		pcbc.Board35x45,
		Board,
	},
}

func TestBoard(t *testing.T) {
	assert.NoError(t, eda.GeneratePCB(testBoard))
}
