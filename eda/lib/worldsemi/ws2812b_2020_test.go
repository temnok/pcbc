// Copyright © 2025 Alex Temnok. All rights reserved.

package worldsemi

import (
	"github.com/stretchr/testify/assert"
	"github.com/temnok/pcbc/eda/pcb"
	"github.com/temnok/pcbc/eda/pcb/config"
	"testing"
)

func TestBoard(t *testing.T) {
	conf := config.Default()
	conf.Width, conf.Height = 4, 3

	assert.NoError(t, pcb.Process(conf, WS2812B_2020))
}
