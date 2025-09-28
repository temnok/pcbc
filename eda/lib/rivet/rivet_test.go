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

				Tracks: path.Paths{
					eda.LinearTrack(path.Point{-4, 0}, path.Point{-1, 0}),
					eda.LinearTrack(path.Point{1, 0}, path.Point{3, 0}),
				},
			},

			{
				Back: true,

				Tracks: path.Paths{
					eda.LinearTrack(path.Point{-1, 0}, path.Point{1, 0}),
				},
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

			rivetPair.CloneY(3, 2).Arrange(transform.Move(0.5, 0)),
			rivetPair.CloneY(2, 2).Arrange(transform.RotateDegrees(180).Move(-0.5, 0)),
		},
	}

	conf := config.Default()
	conf.Width, conf.Height = 14, 14

	assert.NoError(t, pcb.Process(conf, board))
}
