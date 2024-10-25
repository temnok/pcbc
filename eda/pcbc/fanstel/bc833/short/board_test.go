package long

import (
	"github.com/stretchr/testify/assert"
	"temnok/pcbc/eda"
	"temnok/pcbc/eda/lib"
	"temnok/pcbc/eda/pcbc"
	"temnok/pcbc/eda/pcbc/fanstel/bc833"
	"temnok/pcbc/transform"
	"testing"
)

func Test_BC833Short(t *testing.T) {
	assert.NoError(t, eda.GeneratePCB(&lib.Component{
		Components: lib.Components{
			pcbc.Board35x45,
			bc833.ShortBoard.Arrange(transform.Move(0, 10.5)),
			bc833.ShortBoard.Arrange(transform.Rotate(180).Move(0, -10.5)),
		},
	}))
}
