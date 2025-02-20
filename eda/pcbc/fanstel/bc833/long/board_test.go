// Copyright © 2025 Alex Temnok. All rights reserved.

package long

import (
	"github.com/stretchr/testify/assert"
	"temnok/pcbc/eda"
	"temnok/pcbc/eda/pcb"
	"temnok/pcbc/eda/pcbc"
	"temnok/pcbc/eda/pcbc/fanstel/bc833"
	"testing"
)

func Test_BC833(t *testing.T) {
	assert.NoError(t, pcb.Generate(&eda.Component{
		Components: eda.Components{
			pcbc.Board35x45,
			bc833.Board,
		},
	}))
}
