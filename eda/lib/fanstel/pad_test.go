// Copyright © 2025 Alex Temnok. All rights reserved.

package fanstel

import (
	"github.com/stretchr/testify/assert"
	"temnok/pcbc/eda"
	"temnok/pcbc/eda/pcb"
	"temnok/pcbc/eda/pcb/config"
	"temnok/pcbc/path"
	"temnok/pcbc/transform"
	"testing"
)

func TestBoard(t *testing.T) {
	conf := config.Default()
	conf.Width, conf.Height = 13, 18

	x, y := 5.0, 14.3/2

	assert.NoError(t, pcb.Process(conf, &eda.Component{
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

		Nested: eda.Components{
			BC833,
		},
	}))
}
