// Copyright © 2025 Alex Temnok. All rights reserved.

package sop8

import (
	"github.com/stretchr/testify/assert"
	"temnok/pcbc/eda/pcb"
	"temnok/pcbc/eda/pcb/config"
	"testing"
)

func TestBoard(t *testing.T) {
	conf := config.Default()
	conf.Width, conf.Height = 15, 9

	assert.NoError(t, pcb.Process(conf, Board))
}
