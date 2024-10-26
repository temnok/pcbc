package qfn16v2

import (
	"github.com/stretchr/testify/assert"
	"temnok/pcbc/eda"
	"temnok/pcbc/eda/lib"
	"temnok/pcbc/eda/pcbc"
	"testing"
)

func TestBoard(t *testing.T) {
	assert.NoError(t, eda.GeneratePCB(&lib.Component{
		Components: lib.Components{
			pcbc.Board35x45,
			Board.Clone(3, 0, 11),
		},
	}))
}
