package x2v2

import (
	"github.com/stretchr/testify/assert"
	"temnok/pcbc/eda"
	"temnok/pcbc/eda/lib"
	"temnok/pcbc/eda/pcbc"
	"temnok/pcbc/transform"
	"testing"
)

func TestBoard(t *testing.T) {
	assert.NoError(t, eda.GeneratePCB(&lib.Component{
		Components: lib.Components{
			pcbc.Board35x45,
			lib.ComponentGrid(3, 11, 5,
				X2("LED+", "R2V-"),
				X2("LED+", "G3V-"),
				X2("LED+", "B3V-"),
				X2("LED+", "Y2V-"),
				X2("LED+", "W3V-"),
				X2("R ", "50R"),
				X2("R ", "50R"),
				X2("R ", "K10"),
				X2("R ", "K10"),
				X2("R ", "K15"),
				X2("R ", "K15"),
				X2("R ", "K20"),
				X2("R ", "K20"),
				X2("R ", "K25"),
				X2("R ", "K25"),
			).Arrange(transform.Rotate(90)),
		},
	}))
}
