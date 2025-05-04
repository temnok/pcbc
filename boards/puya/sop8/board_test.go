// Copyright © 2025 Alex Temnok. All rights reserved.

package sop8

import (
	"github.com/stretchr/testify/assert"
	"temnok/pcbc/eda/pcb"
	"testing"
)

func TestBoard(t *testing.T) {
	assert.NoError(t, pcb.Generate(Board))
}
