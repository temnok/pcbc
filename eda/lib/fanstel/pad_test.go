package fanstel

import (
	"github.com/stretchr/testify/assert"
	"temnok/pcbc/eda"
	"temnok/pcbc/path"
	"temnok/pcbc/transform"
	"testing"
)

func TestBoard(t *testing.T) {
	x, y := 5.0, 14.3/2

	assert.NoError(t, eda.GeneratePCB(&eda.Component{
		Pads: path.Paths{
			path.Rect(0.5, 2).Apply(transform.Move(x, y)),
			path.Rect(2, 0.5).Apply(transform.Move(x, y)),

			path.Rect(0.5, 2).Apply(transform.Move(x, -y)),
			path.Rect(2, 0.5).Apply(transform.Move(x, -y)),

			path.Rect(0.5, 2).Apply(transform.Move(-x, -y)),
			path.Rect(2, 0.5).Apply(transform.Move(-x, -y)),

			path.Rect(0.5, 2).Apply(transform.Move(-x, y)),
			path.Rect(2, 0.5).Apply(transform.Move(-x, y)),
		},
		Components: eda.Components{
			BC833,
		},
	}))
}
