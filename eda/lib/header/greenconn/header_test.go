package greenconn

import (
	"github.com/stretchr/testify/assert"
	"temnok/pcbc/eda"
	"temnok/pcbc/eda/lib"
	"temnok/pcbc/transform"
	"testing"
)

func TestHeader(t *testing.T) {
	pcb := eda.NewPCB(20, 10)

	pcb.Component(&lib.Component{
		Components: lib.Components{
			CSCC118(7, []string{"P001", "P002", "VDD", "D+", "D-", "GND", "P007"}).
				Arrange(transform.Move(-5, 0)),
			CSCC118(8, []string{"P001", "GND", "VDD", "D+", "D-", "P006", "P008", "P009"}).
				Arrange(transform.Move(5, 0)),
		},
	})

	assert.NoError(t, pcb.SaveFiles("gen/"))
}
