// Copyright © 2025 Alex Temnok. All rights reserved.

package x4

import (
	"github.com/stretchr/testify/assert"
	"temnok/pcbc/eda"
	"temnok/pcbc/eda/pcbc"
	"temnok/pcbc/eda/pcbc/puya/qfn16"
	"temnok/pcbc/transform"
	"testing"
)

func TestBoard(t *testing.T) {
	assert.NoError(t, eda.GeneratePCB(&eda.Component{
		Components: eda.Components{
			pcbc.Perforations34x42,
			qfn16.Board.Arrange(transform.Rotate(90)).Clone(2, 12.5, 0).Clone(2, 0, 17),
		},
	}))
}
