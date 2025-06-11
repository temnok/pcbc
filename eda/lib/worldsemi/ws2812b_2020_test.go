// Copyright © 2025 Alex Temnok. All rights reserved.

package worldsemi

import (
	"github.com/stretchr/testify/assert"
	"temnok/pcbc/eda/pcb"
	"testing"
)

func TestBoard(t *testing.T) {
	assert.NoError(t, pcb.Process(nil, WS2812B_2020))
}
