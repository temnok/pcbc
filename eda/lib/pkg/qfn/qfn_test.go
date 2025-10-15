// Copyright © 2025 Alex Temnok. All rights reserved.

package qfn

import (
	"github.com/stretchr/testify/assert"
	"github.com/temnok/pcbc/eda/pcb"
	"github.com/temnok/pcbc/eda/pcb/config"
	"testing"
)

func TestQFN16(t *testing.T) {
	conf := config.Default()
	conf.Width, conf.Height = 5, 5

	assert.NoError(t, pcb.Process(conf, QFN16G))
}
