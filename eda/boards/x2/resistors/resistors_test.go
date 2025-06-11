// Copyright © 2025 Alex Temnok. All rights reserved.

package resistors

import (
	"github.com/stretchr/testify/assert"
	"temnok/pcbc/eda"
	"temnok/pcbc/eda/boards"
	"temnok/pcbc/eda/boards/x2"
	"temnok/pcbc/eda/pcb"
	"temnok/pcbc/eda/pcb/config"
	"testing"
)

func TestBoard(t *testing.T) {
	conf := config.Default()
	conf.Width, conf.Height = 76, 46

	r := func(label string) *eda.Component {
		return x2.X2_I0402("R ", label)
	}

	assert.NoError(t, pcb.Process(conf, &eda.Component{
		Inner: eda.Components{
			boards.Guides72x42,
			eda.ComponentGrid(6, 11, 5,
				r("33R"),
				r("47R"),
				r("68R"),
				r("33R"),
				r("47R"),
				r("68R"),

				r("K10"),
				r("K15"),
				r("K22"),
				r("K33"),
				r("K47"),
				r("K68"),
				r("K10"),
				r("K15"),
				r("K22"),
				r("K33"),
				r("K47"),
				r("K68"),

				r("1K0"),
				r("1K5"),
				r("2K2"),
				r("3K3"),
				r("4K7"),
				r("6K8"),
				r("1K0"),
				r("1K5"),
				r("2K2"),
				r("3K3"),
				r("4K7"),
				r("6K8"),

				r("10K"),
				r("15K"),
				r("22K"),
				r("33K"),
				r("47K"),
				r("68K"),
				r("10K"),
				r("15K"),
				r("22K"),
				r("33K"),
				r("47K"),
				r("68K"),

				r("M10"),
				r("M15"),
				r("M22"),
				r("M10"),
				r("M15"),
				r("M22"),
			),
		},
	}))
}
