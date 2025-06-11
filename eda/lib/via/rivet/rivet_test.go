// Copyright © 2025 Alex Temnok. All rights reserved.

package rivet

import (
	"github.com/stretchr/testify/assert"
	"temnok/pcbc/eda"
	"temnok/pcbc/eda/pcb"
	"temnok/pcbc/eda/pcb/config"
	"temnok/pcbc/path"
	"temnok/pcbc/transform"
	"temnok/pcbc/util"
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

	rivets := &eda.Component{
		Components: eda.Components{
			Rivet06mm.Arrange(transform.Move(-1, 0)),
			Rivet06mm.Arrange(transform.Move(1, 0)),
		},
	}

	rivetPairTop := &eda.Component{
		Components: eda.Components{
			rivets,

			{
				Pads: path.Paths{
					path.Circle(1).Transform(transform.Move(-3, 0)),
					path.Circle(1).Transform(transform.Move(3, 0)),
				},
			},
		},

		Tracks: eda.Tracks(
			eda.Track{{-3, 0}, {-1, 0}},
			eda.Track{{1, 0}, {3, 0}},
		),
	}

	rivetPairBottom := &eda.Component{
		Components: eda.Components{
			rivets,
		},

		Tracks: eda.Tracks(
			eda.Track{{-1, 0}, {1, 0}},
		),
	}

	top := &eda.Component{
		Components: eda.Components{
			blank,

			rivetPairTop.Arrange(transform.Move(0.5, 3)),
			rivetPairTop.Arrange(transform.Move(-0.5, 2)),
			rivetPairTop.Arrange(transform.Move(0.5, 1)),
			rivetPairTop.Arrange(transform.Move(-0.5, 0)),
			rivetPairTop.Arrange(transform.Move(0.5, -1)),
			rivetPairTop.Arrange(transform.Move(-0.5, -2)),
			rivetPairTop.Arrange(transform.Move(0.5, -3)),
		},
	}

	bottom := &eda.Component{
		Transform: transform.MirrorX(),

		Components: eda.Components{
			blank,

			rivetPairBottom.Arrange(transform.Move(0.5, 3)),
			rivetPairBottom.Arrange(transform.Move(-0.5, 2)),
			rivetPairBottom.Arrange(transform.Move(0.5, 1)),
			rivetPairBottom.Arrange(transform.Move(-0.5, 0)),
			rivetPairBottom.Arrange(transform.Move(0.5, -1)),
			rivetPairBottom.Arrange(transform.Move(-0.5, -2)),
			rivetPairBottom.Arrange(transform.Move(0.5, -3)),
		},
	}

	assert.NoError(t, util.RunConcurrently(
		func() error {
			conf := config.Default()
			conf.SavePath = "out/1-"
			return pcb.Process(conf, top)
		},

		func() error {
			conf := config.Default()
			conf.SavePath = "out/2-"
			return pcb.Process(conf, bottom)
		},
	))
}
