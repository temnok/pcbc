// Copyright © 2025 Alex Temnok. All rights reserved.

package long

import (
	"github.com/stretchr/testify/assert"
	"temnok/pcbc/eda"
	"temnok/pcbc/eda/boards"
	"temnok/pcbc/eda/boards/fanstel/bc833"
	"temnok/pcbc/eda/pcb"
	"temnok/pcbc/transform"
	"testing"
)

func Test_BC833Short(t *testing.T) {
	assert.NoError(t, pcb.ProcessWithDefaultConfig(&eda.Component{
		Components: eda.Components{
			boards.Board35x45,
			bc833.ShortBoard.Arrange(transform.Move(0, 10.5)),
			bc833.ShortBoard.Arrange(transform.RotateDegrees(180).Move(0, -10.5)),
		},
	}))
}
