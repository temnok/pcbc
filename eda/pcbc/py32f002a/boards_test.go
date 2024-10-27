package py32f002a

import (
	"github.com/stretchr/testify/assert"
	"temnok/pcbc/eda"
	"temnok/pcbc/eda/pcbc"
	"temnok/pcbc/eda/pcbc/py32f002a/qfn16tiny"
	"temnok/pcbc/eda/pcbc/py32f002a/sop8tiny"
	"temnok/pcbc/transform"
	"testing"
)

func TestBoards(t *testing.T) {
	assert.NoError(t, eda.GeneratePCB(&eda.Component{
		Components: eda.Components{
			pcbc.Board35x45,
			sop8tiny.Board.Arrange(transform.Rotate(-90)).Clone(3, 9.5, 0).Arrange(transform.Move(0, 8.5)),
			qfn16tiny.Board.Arrange(transform.Rotate(90)).Clone(2, 11, 0).Arrange(transform.Move(0, -8.5)),
		},
	}))
}
