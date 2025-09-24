// Copyright © 2025 Alex Temnok. All rights reserved.

package e73

import (
	"github.com/stretchr/testify/assert"
	"temnok/pcbc/eda/pcb"
	"temnok/pcbc/eda/pcb/config"
	"testing"
)

func TestBoard(t *testing.T) {
	conf := config.Default()
	conf.Width = 30
	conf.Height = 20

	assert.NoError(t, pcb.Process(conf, Board_nRF52840))
}
