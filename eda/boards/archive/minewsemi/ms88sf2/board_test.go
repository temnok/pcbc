// Copyright © 2025 Alex Temnok. All rights reserved.

package ms88sf2

import (
	"github.com/stretchr/testify/assert"
	"temnok/pcbc/eda"
	"temnok/pcbc/eda/boards"
	"temnok/pcbc/eda/pcb"
	"temnok/pcbc/eda/pcb/config"
	"testing"
)

func TestBoard(t *testing.T) {
	conf := config.Default()
	conf.Width, conf.Height = 36, 46

	assert.NoError(t, pcb.Process(conf, &eda.Component{
		Inner: eda.Components{
			boards.Board35x45,
			Board_nRF52840,
		},
	}))
}
