package x2tiny

import (
	"github.com/stretchr/testify/assert"
	"temnok/pcbc/eda"
	"temnok/pcbc/eda/lib"
	"temnok/pcbc/eda/pcbc"
	"testing"
)

func TestDemo(t *testing.T) {
	assert.NoError(t, eda.GeneratePCB(&lib.Component{
		Components: lib.Components{
			pcbc.Board35x45,
			X2Base,
		},
	}))
}
