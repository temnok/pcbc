// Copyright © 2025 Alex Temnok. All rights reserved.

package rivet

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
	hole := path.Circle(1.45)

	rivetPair := &eda.Component{
		Nested: eda.Components{
			Rivet.Arrange(transform.Move(-1, 0)),
			Rivet.Arrange(transform.Move(1, 0)),

			{
				Pads: path.Paths{
					path.RoundRect(1, 0.55, 0.15).Transform(transform.Move(-4, 0)),
					path.RoundRect(1, 0.55, 0.15).Transform(transform.Move(3, 0)),
				},

				Tracks: eda.Tracks(
					eda.Track{{-4, 0}, {-1, 0}},
					eda.Track{{1, 0}, {3, 0}},
				),
			},

			{
				Back: true,

				Tracks: eda.Tracks(
					eda.Track{{-1, 0}, {1, 0}},
				),
			},
		},
	}

	board := &eda.Component{
		Nested: eda.Components{
			{
				CutsOuter: true,

				Cuts: path.Paths{
					path.RoundRect(13, 13, 1.5),
				},
			},

			{
				CutsFully: true,

				Cuts: path.Paths{
					hole.Transform(transform.Move(-5, -5)),
					hole.Transform(transform.Move(-5, 5)),
					hole.Transform(transform.Move(5, -5)),
					hole.Transform(transform.Move(5, 5)),
				},
			},

			rivetPair.Clone(3, 0, 2).Arrange(transform.Move(0.5, 0)),
			rivetPair.Clone(2, 0, 2).Arrange(transform.RotateDegrees(180).Move(-0.5, 0)),
		},
	}

	conf := config.Default()
	conf.Width, conf.Height = 14, 14

	assert.NoError(t, pcb.Process(conf, board))
}
