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
	perforation := path.Circle(1)

	perforations := &eda.Component{
		Components: eda.Components{
			{
				Perforations: path.Paths{
					perforation.Apply(transform.Move(-10, 4)),
					perforation.Apply(transform.Move(-10, -6)),
					perforation.Apply(transform.Move(0, 6)),
					perforation.Apply(transform.Move(0, -6)),
					perforation.Apply(transform.Move(10, 6)),
					perforation.Apply(transform.Move(10, -6)),
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
		TrackWidth: 0.15,
		Tracks: eda.Tracks(
			eda.Track{{X: -2.4, Y: -0.25}}.DX(3.5),
			eda.Track{{X: -1.2, Y: 0.25}}.DX(3.5),
		).Clone(5, 0, 1),
	}

	top := pcb.New(&eda.Component{
		Holes: vias,

		Components: eda.Components{
			perforations,

			contacts,

			pads.Clone(2, 4, 0),

			tracks.Clone(2, 8, 0),
		},
	})
	top.Width += 4
	top.Height += 4
	top.SaveEtchOverride = top.SaveEtchPI
	top.SavePath = "out/1-"

	bottom := pcb.New(&eda.Component{
		Pads: bottomPads,

		Components: eda.Components{
			perforations,

			tracks,
		},
	})
	bottom.Width += 4
	bottom.Height += 4
	bottom.SaveEtchOverride = bottom.SaveEtchPI
	bottom.SavePath = "out/2-"

	assert.NoError(t, top.Process().SaveFiles())
	assert.NoError(t, bottom.Process().SaveFiles())
}
