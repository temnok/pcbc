// Copyright © 2025 Alex Temnok. All rights reserved.

package fanstel

import (
	"github.com/stretchr/testify/assert"
	"temnok/pcbc/eda"
	"temnok/pcbc/eda/pcb"
	"temnok/pcbc/path"
	"temnok/pcbc/transform"
	"testing"
)

func TestBoard(t *testing.T) {
	x, y := 5.0, 14.3/2

	assert.NoError(t, pcb.ProcessWithDefaultConfig(&eda.Component{
		Pads: path.Paths{
			path.Rect(0.5, 2).Transform(transform.Move(x, y)),
			path.Rect(2, 0.5).Transform(transform.Move(x, y)),

			path.Rect(0.5, 2).Transform(transform.Move(x, -y)),
			path.Rect(2, 0.5).Transform(transform.Move(x, -y)),

			path.Rect(0.5, 2).Transform(transform.Move(-x, -y)),
			path.Rect(2, 0.5).Transform(transform.Move(-x, -y)),

			path.Rect(0.5, 2).Transform(transform.Move(-x, y)),
			path.Rect(2, 0.5).Transform(transform.Move(-x, y)),
		},

		Components: eda.Components{
			BC833,
		},
	}))
}
