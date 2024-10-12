package x2

import (
	"github.com/stretchr/testify/assert"
	"temnok/pcbc/eda"
	"temnok/pcbc/eda/lib"
	"temnok/pcbc/eda/pcbc"
	"temnok/pcbc/geom"
	"testing"
)

var demoBoard = &lib.Component{
	//Clears: path.Paths{path.Rect(36, 46)},
	Components: lib.Components{
		pcbc.Board35x45,
		{
			Transform: geom.RotateD(-90),
			Components: lib.ComponentsGrid(3, 5, 10.5, 5.5,
				X0603WithGround("LED+", "RED 2V", "20mA"),
				X0603WithGround("LED+", "GRN 3V", "20mA"),
				X0603WithGround("LED+", "BLU 3V", "20mA"),
				X0603WithGround("LED+", "YEL 2V", "20mA"),
				X0603WithGround("LED+", "WHT 3V", "20mA"),
				X0603("75R", "75 Ohm", "Resistor"),
				X0603("75R", "75 Ohm", "Resistor"),
				X0603("K10", "100 Ohm", "Resistor"),
				X0603("K10", "100 Ohm", "Resistor"),
				X0603("K15", "150 Ohm", "Resistor"),
				X0603("K15", "150 Ohm", "Resistor"),
				X0603("K20", "200 Ohm", "Resistor"),
				X0603("K20", "200 Ohm", "Resistor"),
				X0603("K30", "300 Ohm", "Resistor"),
				X0603("K30", "300 Ohm", "Resistor"),
			),
		},
	},
}

func TestDemo(t *testing.T) {
	pcb := eda.NewPCB(36, 46)
	pcb.Component(demoBoard)

	assert.NoError(t, pcb.SaveFiles("out/"))
}
