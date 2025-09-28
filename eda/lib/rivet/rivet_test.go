// Copyright © 2025 Alex Temnok. All rights reserved.

package rivet

import (
	"github.com/stretchr/testify/assert"
	"temnok/pcbc/eda"
	"temnok/pcbc/eda/boards"
	"temnok/pcbc/eda/pcb"
	"temnok/pcbc/eda/pcb/config"
	"temnok/pcbc/path"
	"temnok/pcbc/transform"
	"testing"
)

func TestBoard(t *testing.T) {
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
					path.RoundRect(11, 11, 1.5),
				},
			},

			boards.AlignHole.CloneX(2, 8).CloneY(2, 8),

			rivetPair.CloneY(3, 2).Arrange(transform.Move(0.5, 0)),
			rivetPair.CloneY(2, 2).Arrange(transform.RotateDegrees(180).Move(-0.5, 0)),
		},
	}

	conf := config.Default()
	conf.Width, conf.Height = 12, 12

	assert.NoError(t, pcb.Process(conf, board))
}
