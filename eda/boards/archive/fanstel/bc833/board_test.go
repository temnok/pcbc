// Copyright © 2025 Alex Temnok. All rights reserved.

package bc833

import (
	"github.com/stretchr/testify/assert"
	"temnok/pcbc/eda"
	"temnok/pcbc/eda/pcb"
	"temnok/pcbc/transform"
	"testing"
)

func Test_BC833(t *testing.T) {
	assert.NoError(t, pcb.Process(nil, &eda.Component{
		Inner: eda.Components{
			Board.Arrange(transform.Move(0, 10.5)),
			Board.Arrange(transform.RotateDegrees(180).Move(0, -10.5)),
		},
	}))
}
