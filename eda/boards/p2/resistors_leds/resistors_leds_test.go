// Copyright © 2025 Alex Temnok. All rights reserved.

package resistors

import (
	"github.com/stretchr/testify/assert"
	"temnok/pcbc/eda"
	"temnok/pcbc/eda/boards/p2"
	"temnok/pcbc/eda/pcb"
	"temnok/pcbc/eda/pcb/config"
	"testing"
)

func TestBoard(t *testing.T) {
	conf := config.Default()
	conf.Width, conf.Height = 76, 46

	assert.NoError(t, pcb.Process(conf, &eda.Component{
		Nested: eda.Components{
			eda.ComponentGrid(6, 11, 5,
				p2.P2_I0402("R ", "1R0"),
				p2.P2_I0402("R ", "1R5"),
				p2.P2_I0402("R ", "2R2"),
				p2.P2_I0402("R ", "3R3"),
				p2.P2_I0402("R ", "4R7"),
				p2.P2_I0402("R ", "6R8"),

				p2.P2_I0402("R ", "10R"),
				p2.P2_I0402("R ", "15R"),
				p2.P2_I0402("R ", "22R"),
				p2.P2_I0402("R ", "33R"),
				p2.P2_I0402("R ", "47R"),
				p2.P2_I0402("R ", "68R"),

				p2.P2_I0402("R ", "K10"),
				p2.P2_I0402("R ", "K15"),
				p2.P2_I0402("R ", "K22"),
				p2.P2_I0402("R ", "K33"),
				p2.P2_I0402("R ", "K47"),
				p2.P2_I0402("R ", "K68"),

				p2.P2_I0402("R ", "1K0"),
				p2.P2_I0402("R ", "1K5"),
				p2.P2_I0402("R ", "2K2"),
				p2.P2_I0402("R ", "3K3"),
				p2.P2_I0402("R ", "4K7"),
				p2.P2_I0402("R ", "6K8"),

				p2.P2_I0402("R ", "10K"),
				p2.P2_I0402("R ", "15K"),
				p2.P2_I0402("R ", "22K"),
				p2.P2_I0402("R ", "33K"),
				p2.P2_I0402("R ", "47K"),
				p2.P2_I0402("R ", "68K"),

				p2.P2_I0402("R ", "M10"),
				p2.P2_I0402("R ", "M15"),
				p2.P2_I0402("R ", "M22"),
				p2.P2_I0402("R ", "M33"),
				p2.P2_I0402("R ", "M47"),
				p2.P2_I0402("R ", "M68"),

				p2.P2_I0402("R ", "1M0"),
				p2.P2_I0402("R ", "1M5"),
				p2.P2_I0402("R ", "2M2"),
				p2.P2_I0402("R ", "3M3"),
				p2.P2_I0402("R ", "4M7"),
				p2.P2_I0402("R ", "6M8"),

				p2.P2_I0402("R ", "10M"),

				p2.P2_I0402("+LR", "-2V"),
				p2.P2_I0402("+LG", "-3V"),
				p2.P2_I0402("+LB", "-3V"),
				p2.P2_I0402("+LY", "-2V"),
				p2.P2_I0402("+LW", "-3V"),
			),
		},
	}))
}
