// Copyright © 2025 Alex Temnok. All rights reserved.

package x4

import (
	"github.com/stretchr/testify/assert"
	"temnok/pcbc/eda"
	"temnok/pcbc/eda/boards"
	"temnok/pcbc/eda/boards/puya/qfn16"
	"temnok/pcbc/eda/pcb"
	"temnok/pcbc/eda/pcb/config"
	"temnok/pcbc/transform"
	"testing"
)

func TestBoard(t *testing.T) {
	conf := config.Default()
	conf.Width, conf.Height = 38, 46

	assert.NoError(t, pcb.Process(conf, &eda.Component{
		Inner: eda.Components{
			boards.Guides34x42,
			qfn16.Board.Arrange(transform.RotateDegrees(90)).Clone(2, 12.5, 0).Clone(2, 0, 17),
		},
	}))
}
