// Copyright © 2025 Alex Temnok. All rights reserved.

package resistors

import (
	"github.com/stretchr/testify/assert"
	"temnok/pcbc/eda"
	"temnok/pcbc/eda/boards"
	"temnok/pcbc/eda/boards/x2"
	"temnok/pcbc/eda/pcb"
	"testing"
)

func TestBoard(t *testing.T) {
	assert.NoError(t, pcb.ProcessWithDefaultConfig(&eda.Component{
		Components: eda.Components{
			boards.Guides72x42,
			eda.ComponentGrid(6, 11, 5,
				x2.X2_I0603("R ", "1R0"),
				x2.X2_I0603("R ", "1R5"),
				x2.X2_I0603("R ", "2R2"),
				x2.X2_I0603("R ", "3R3"),
				x2.X2_I0603("R ", "4R7"),
				x2.X2_I0603("R ", "6R8"),

				x2.X2_I0603("R ", "10R"),
				x2.X2_I0603("R ", "15R"),
				x2.X2_I0603("R ", "22R"),
				x2.X2_I0603("R ", "33R"),
				x2.X2_I0603("R ", "47R"),
				x2.X2_I0603("R ", "68R"),

				x2.X2_I0603("R ", "K10"),
				x2.X2_I0603("R ", "K15"),
				x2.X2_I0603("R ", "K22"),
				x2.X2_I0603("R ", "K33"),
				x2.X2_I0603("R ", "K47"),
				x2.X2_I0603("R ", "K68"),

				x2.X2_I0603("R ", "1K0"),
				x2.X2_I0603("R ", "1K5"),
				x2.X2_I0603("R ", "2K2"),
				x2.X2_I0603("R ", "3K3"),
				x2.X2_I0603("R ", "4K7"),
				x2.X2_I0603("R ", "6K8"),

				x2.X2_I0603("R ", "10K"),
				x2.X2_I0603("R ", "15K"),
				x2.X2_I0603("R ", "22K"),
				x2.X2_I0603("R ", "33K"),
				x2.X2_I0603("R ", "47K"),
				x2.X2_I0603("R ", "68K"),

				x2.X2_I0603("R ", "M10"),
				x2.X2_I0603("R ", "M15"),
				x2.X2_I0603("R ", "M22"),
				x2.X2_I0603("R ", "M33"),
				x2.X2_I0603("R ", "M47"),
				x2.X2_I0603("R ", "M68"),

				x2.X2_I0603("R ", "1M0"),
				x2.X2_I0603("R ", "1M5"),
				x2.X2_I0603("R ", "2M2"),
				x2.X2_I0603("R ", "3M3"),
				x2.X2_I0603("R ", "4M7"),
				x2.X2_I0603("R ", "6M8"),

				x2.X2_I0603("R ", "10M"),

				x2.X2_I0603("LD+", "R2V"),
				x2.X2_I0603("LD+", "G3V"),
				x2.X2_I0603("LD+", "B3V"),
				x2.X2_I0603("LD+", "Y2V"),
				x2.X2_I0603("LD+", "W3V"),
			),
		},
	}))
}
