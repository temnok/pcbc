// Copyright © 2025 Alex Temnok. All rights reserved.

package ts026a

import (
	"github.com/stretchr/testify/assert"
	"temnok/pcbc/eda"
	"testing"
)

func TestBoard(t *testing.T) {
	assert.NoError(t, eda.GeneratePCB(Board))
}
