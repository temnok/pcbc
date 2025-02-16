package via

import (
	"github.com/stretchr/testify/assert"
	"temnok/pcbc/eda"
	"temnok/pcbc/eda/pcbc"
	"temnok/pcbc/path"
	"temnok/pcbc/transform"
	"testing"
)

func TestBoard(t *testing.T) {
	via := &eda.Component{
		Holes: path.Circle(0.6),
		Pads:  path.Circle(1.2),
	}

	board := &eda.Component{
		Cuts: path.RoundRect(30, 40, 1),

		Components: eda.Components{
			via.Arrange(transform.Move(-12.5, 17.5)),
			via.Clone(5, 5, 0).Clone(7, 0, 5),
		},
	}

	top := &eda.Component{
		Tracks: eda.Tracks(
			eda.Track{{-10, 15}}.DY(-5),
			eda.Track{{-10, 5}}.DY(-5),
			eda.Track{{-10, -5}}.DY(-5),
			eda.Track{{-10, -15}}.DX(5),
			eda.Track{{0, -15}}.DX(5),
			eda.Track{{10, -15}}.DY(5),
			eda.Track{{10, -5}}.DY(5),
			eda.Track{{10, 5}}.DY(5),
			eda.Track{{10, 15}}.DX(-5),
			eda.Track{{0, 15}}.DX(-5),
			eda.Track{{-5, 10}}.DY(-5),
			eda.Track{{-5, 0}}.DY(-5),
			eda.Track{{-5, -10}}.DX(5),
			eda.Track{{5, -10}}.DY(5),
			eda.Track{{5, 0}}.DY(5),
			eda.Track{{5, 10}}.DX(-5),
			eda.Track{{0, 5}}.DY(-5),
		),

		Components: eda.Components{
			board,
		},
	}

	bottom := &eda.Component{
		Tracks: eda.Tracks(
			eda.Track{{10, 10}}.DY(-5),
			eda.Track{{10, 0}}.DY(-5),
			eda.Track{{10, -10}}.DY(-5),
			eda.Track{{5, -15}}.DX(-5),
			eda.Track{{-5, -15}}.DX(-5),
			eda.Track{{-10, -10}}.DY(5),
			eda.Track{{-10, 0}}.DY(5),
			eda.Track{{-10, 10}}.DY(5),
			eda.Track{{-5, 15}}.DX(5),
			eda.Track{{5, 15}}.DY(-5),
			eda.Track{{5, 5}}.DY(-5),
			eda.Track{{5, -5}}.DY(-5),
			eda.Track{{0, -10}}.DX(-5),
			eda.Track{{-5, -5}}.DY(5),
			eda.Track{{-5, 5}}.DY(5),
			eda.Track{{0, 10}}.DY(-5),
			eda.Track{{0, 0}}.DY(-5),
		),

		Components: eda.Components{
			board.Arrange(transform.Scale(-1, 1)),
		},
	}

	assert.NoError(t, eda.GeneratePCB(&eda.Component{
		Components: eda.Components{
			pcbc.Base72x42,
			top.Arrange(transform.Move(-16, 0)),
			bottom.Arrange(transform.Move(16, 0)),
		},
	}))
}
