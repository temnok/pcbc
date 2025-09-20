// Copyright © 2025 Alex Temnok. All rights reserved.

package lir1254

import (
	"github.com/stretchr/testify/assert"
	"temnok/pcbc/eda/pcb"
	"temnok/pcbc/eda/pcb/config"
	"temnok/pcbc/transform"
	"testing"
)

func TestBoard(t *testing.T) {
	conf := config.Default()
	conf.Width, conf.Height = 21, 19

	conf.SavePath = "out/{}-"

	assert.NoError(t, pcb.Process(conf,
		Board.InLayer(1),
		Board.InLayer(2).Arrange(transform.MirrorX()),
	))
}
