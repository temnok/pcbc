package via_pi

import (
	"github.com/stretchr/testify/assert"
	"temnok/pcbc/eda"
	"temnok/pcbc/eda/pcb"
	"temnok/pcbc/path"
	"temnok/pcbc/transform"
	"testing"
)

func TestBoard(t *testing.T) {
	director := path.Circle(1)

	directors := &eda.Component{
		Components: eda.Components{
			{
				Perforations: path.Paths{
					director.Apply(transform.Move(-10, 5)),
					director.Apply(transform.Move(-10, -6)),
					director.Apply(transform.Move(10, 6)),
					director.Apply(transform.Move(10, -6)),
				},
			},
		},
	}

	viaHole := &eda.Component{
		Holes: path.Paths{path.Rect(0.4, 0.4)},
	}

	viaHoles := &eda.Component{
		Components: eda.Components{
			viaHole.Clone(2, 1, 0.5).Clone(5, 0, 1),
		},
	}

	pad := &eda.Component{
		Pads: path.Paths{path.RoundRect(0.8, 0.4, 0.1)},
	}

	pads := &eda.Component{
		Components: eda.Components{
			pad.Clone(2, 1, 0.5).Clone(5, 0, 1),
		},
	}

	tracks := &eda.Component{
		TrackWidth: 0.15,
		Tracks: eda.Tracks(
			eda.Track{{-2.7, -0.25}}.DX(4),
			eda.Track{{-1.3, 0.25}}.DX(4),
		).Clone(5, 0, 1),
	}

	top := pcb.Process(&eda.Component{
		Components: eda.Components{
			directors,

			pads.Clone(4, 4, 0),

			viaHoles.Clone(2, 3.6, 0),

			tracks.Clone(2, 8, 0),
		},
	})
	top.SavePath = "out/1-"

	bottom := pcb.Process(&eda.Component{
		Components: eda.Components{
			directors,

			pads.Clone(2, 4, 0),

			tracks,
		},
	})
	bottom.SavePath = "out/2-"

	assert.NoError(t, top.SaveFiles())
	assert.NoError(t, bottom.SaveFiles())
}
