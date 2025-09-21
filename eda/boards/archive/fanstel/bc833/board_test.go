// Copyright © 2025 Alex Temnok. All rights reserved.

package bc833

import (
	"github.com/stretchr/testify/assert"
	"temnok/pcbc/eda"
	"temnok/pcbc/eda/pcb"
	"temnok/pcbc/eda/pcb/config"
	"temnok/pcbc/transform"
	"testing"
)

func Test_BC833(t *testing.T) {
	conf := config.Default()
	conf.Width, conf.Height = 22, 51

	assert.NoError(t, pcb.Process(conf, &eda.Component{
		Nested: eda.Components{
			Board.Arrange(transform.Move(0, 10.5)),
			Board.Arrange(transform.RotateDegrees(180).Move(0, -10.5)),
		},
	}))
}
