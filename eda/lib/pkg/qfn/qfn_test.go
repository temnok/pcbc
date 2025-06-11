// Copyright © 2025 Alex Temnok. All rights reserved.

package qfn

import (
	"github.com/stretchr/testify/assert"
	"temnok/pcbc/eda/pcb"
	"temnok/pcbc/eda/pcb/config"
	"testing"
)

func TestQFN16(t *testing.T) {
	conf := config.Default()
	conf.Width, conf.Height = 5, 5

	assert.NoError(t, pcb.Process(conf, QFN16G))
}
