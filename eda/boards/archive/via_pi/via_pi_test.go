// Copyright © 2025 Alex Temnok. All rights reserved.

package via_pi

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
	hole := path.Circle(1)

	holes := &eda.Component{
		Inner: eda.Components{
			{
				Cuts: path.Paths{
					hole.Transform(transform.Move(-10, 4)),
					hole.Transform(transform.Move(-10, -6)),
					hole.Transform(transform.Move(0, 6)),
					hole.Transform(transform.Move(0, -6)),
					hole.Transform(transform.Move(10, 6)),
					hole.Transform(transform.Move(10, -6)),
				},
			},
		},
	}

	contacts := &eda.Component{
		Pads: path.Paths{path.RoundRect(2, 0.3, 0.1)}.Clone(10, 0, 0.5).Clone(2, 12, 0),
	}

	vias := path.Paths{path.Rect(0.4, 0.4)}.Clone(2, 1, 0.5).Clone(5, 0, 1).Clone(2, 3.6, 0)
	bottomPads := path.Paths{path.Rect(0.5, 0.5)}.Clone(2, 1, 0.5).Clone(5, 0, 1).Clone(2, 3.6, 0)

	pads := &eda.Component{
		Pads: path.Paths{path.RoundRect(0.8, 0.4, 0.1)}.Clone(2, 1, 0.5).Clone(5, 0, 1),
	}

	tracks := &eda.Component{
		TracksWidth: 0.15,
		Tracks: eda.Tracks(
			eda.Track{{X: -2.4, Y: -0.25}}.DX(3.5),
			eda.Track{{X: -1.2, Y: 0.25}}.DX(3.5),
		).Clone(5, 0, 1),
	}

	top := &eda.Component{
		Pads: vias,

		Inner: eda.Components{
			holes,

			contacts,

			pads.Clone(2, 4, 0),

			tracks.Clone(2, 8, 0),
		},
	}

	topConfig := config.Default()
	topConfig.Width, topConfig.Height = 22, 14
	topConfig.SavePath = "out/1-"

	bottom := &eda.Component{
		Pads: bottomPads,

		Inner: eda.Components{
			holes,

			tracks,
		},
	}

	bottomConfig := config.Default()
	bottomConfig.Width, bottomConfig.Height = 22, 14
	bottomConfig.SavePath = "out/2-"

	assert.NoError(t, pcb.Process(topConfig, top))
	assert.NoError(t, pcb.Process(bottomConfig, bottom))
}
