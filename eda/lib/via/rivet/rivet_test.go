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

	hole := path.Circle(1.4)

	blank := &eda.Component{
		OuterCut: true,

		Cuts: path.Paths{
			path.RoundRect(13, 13, 1.4),

			hole.Transform(transform.Move(-5, -5)),
			hole.Transform(transform.Move(-5, 5)),
			hole.Transform(transform.Move(5, -5)),
			hole.Transform(transform.Move(5, 5)),
		},
	}

	rivetPair := &eda.Component{
		Components: eda.Components{
			Rivet06mm_Layers12.Arrange(transform.Move(-1, 0)),
			Rivet06mm_Layers12.Arrange(transform.Move(1, 0)),

			{
				Layer: 1,

				Pads: path.Paths{
					path.Circle(1).Transform(transform.Move(-3, 0)),
					path.Circle(1).Transform(transform.Move(3, 0)),
				},

				Tracks: eda.Tracks(
					eda.Track{{-3, 0}, {-1, 0}},
					eda.Track{{1, 0}, {3, 0}},
				),
			},

			{
				Layer: 2,

				Tracks: eda.Tracks(
					eda.Track{{-1, 0}, {1, 0}},
				),
			},
		},
	}

	board := &eda.Component{
		Components: eda.Components{
			blank,

			rivetPair.Clone(4, 0, 2).Arrange(transform.Move(0.5, 0)),
			rivetPair.Clone(3, 0, 2).Arrange(transform.Move(-0.5, 0)),
		},
	}

	conf := config.Default()
	conf.SavePath = "out/{}-"

	assert.NoError(t, pcb.Process(conf,
		board.WithLayer(1),
		board.WithLayer(2).Arrange(transform.MirrorX()),
	))
}
