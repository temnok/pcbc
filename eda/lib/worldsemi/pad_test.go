// Copyright © 2025 Alex Temnok. All rights reserved.

package worldsemi

import (
	"github.com/stretchr/testify/assert"
	"temnok/pcbc/eda"
	"testing"
)

func TestBoard(t *testing.T) {
	assert.NoError(t, eda.ProcessPCB(WS2812B_2020).SaveOverview())
}
