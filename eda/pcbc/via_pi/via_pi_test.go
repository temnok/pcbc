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
					director.Apply(transform.Move(-10, 4)),
					director.Apply(transform.Move(-10, -6)),
					director.Apply(transform.Move(0, 6)),
					director.Apply(transform.Move(0, -6)),
					director.Apply(transform.Move(10, 6)),
					director.Apply(transform.Move(10, -6)),
				},
			},
		},
	}

	contacts := &eda.Component{
		Pads: path.Paths{path.RoundRect(2, 0.3, 0.1)}.Clone(10, 0, 0.5).Clone(2, 12, 0),
	}

	vias := path.Paths{path.Rect(0.4, 0.4)}.Clone(2, 1, 0.5).Clone(5, 0, 1).Clone(2, 3.6, 0)

	pads := &eda.Component{
		Pads: path.Paths{path.RoundRect(0.8, 0.4, 0.1)}.Clone(2, 1, 0.5).Clone(5, 0, 1),
	}

	tracks := &eda.Component{
		TrackWidth: 0.15,
		Tracks: eda.Tracks(
			eda.Track{{-2.4, -0.25}}.DX(3.5),
			eda.Track{{-1.2, 0.25}}.DX(3.5),
		).Clone(5, 0, 1),
	}

	top := pcb.Process(&eda.Component{
		Holes: vias,

		Components: eda.Components{
			directors,

			contacts,

			pads.Clone(2, 4, 0),

			tracks.Clone(2, 8, 0),
		},
	})
	top.SavePath = "out/1-"

	bottom := pcb.Process(&eda.Component{
		Pads: vias,

		Components: eda.Components{
			directors,

			tracks,
		},
	})
	bottom.SavePath = "out/2-"

	assert.NoError(t, top.SaveFiles())
	assert.NoError(t, bottom.SaveFiles())
}
