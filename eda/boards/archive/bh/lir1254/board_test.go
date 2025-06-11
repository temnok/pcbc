// Copyright © 2025 Alex Temnok. All rights reserved.

package lir1254

import (
	"github.com/stretchr/testify/assert"
	"temnok/pcbc/eda"
	"temnok/pcbc/eda/boards"
	"temnok/pcbc/eda/pcb"
	"temnok/pcbc/eda/pcb/config"
	"testing"
)

var testBoard = &eda.Component{
	Inner: eda.Components{
		boards.Board35x45,
		Board,
	},
}

func TestBoard(t *testing.T) {
	conf := config.Default()
	conf.Width, conf.Height = 36, 46

	assert.NoError(t, pcb.Process(conf, testBoard))
}
