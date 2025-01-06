// Copyright © 2025 Alex Temnok. All rights reserved.

package greenconn

import (
	"github.com/stretchr/testify/assert"
	"temnok/pcbc/eda"
	"temnok/pcbc/transform"
	"testing"
)

func TestHeader(t *testing.T) {
	assert.NoError(t, eda.NewPCB(&eda.Component{
		Components: eda.Components{
			CSCC118(7, false, []string{"P001", "P002", "VDD", "D+", "D-", "GND", "P007"}).
				Arrange(transform.Move(-5, 0)),
			CSCC118(8, false, []string{"P001", "GND", "VDD", "D+", "D-", "P006", "P008", "P009"}).
				Arrange(transform.Move(5, 0)),
		},
	}).SaveOverview())
}
