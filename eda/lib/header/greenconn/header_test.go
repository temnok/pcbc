// Copyright © 2025 Alex Temnok. All rights reserved.

package greenconn

import (
	"github.com/stretchr/testify/assert"
	"temnok/pcbc/eda"
	"temnok/pcbc/eda/pcb"
	"temnok/pcbc/eda/pcb/config"
	"temnok/pcbc/transform"
	"testing"
)

func TestHeader(t *testing.T) {
	conf := config.Default()
	conf.Width, conf.Height = 17, 10

	assert.NoError(t, pcb.Process(conf, &eda.Component{
		Inner: eda.Components{
			CSCC118(7, false, []string{"P001", "P002", "VDD", "D+", "D-", "GND", "P007"}).
				Arrange(transform.Move(-5, 0)),
			CSCC118(8, false, []string{"P001", "GND", "VDD", "D+", "D-", "P006", "P008", "P009"}).
				Arrange(transform.Move(5, 0)),
		},
	}))
}
