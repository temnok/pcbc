// Copyright © 2025 Alex Temnok. All rights reserved.

package puya

import (
	"github.com/stretchr/testify/assert"
	"temnok/pcbc/boards"
	"temnok/pcbc/boards/puya/qfn16"
	"temnok/pcbc/boards/puya/sop8"
	"temnok/pcbc/eda"
	"temnok/pcbc/eda/pcb"
	"temnok/pcbc/transform"
	"testing"
)

func TestBoards(t *testing.T) {
	assert.NoError(t, pcb.Generate(&eda.Component{
		Components: eda.Components{
			boards.Board35x45,
			sop8.Board.Arrange(transform.RotateDegrees(-90)).Clone(3, 9.5, 0).Arrange(transform.Move(0, 8.5)),
			qfn16.Board.Arrange(transform.RotateDegrees(90)).Clone(2, 11, 0).Arrange(transform.Move(0, -8.5)),
		},
	}))
}
