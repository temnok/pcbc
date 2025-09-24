// Copyright © 2025 Alex Temnok. All rights reserved.

package boards

import (
	"temnok/pcbc/eda"
	"temnok/pcbc/path"
)

var (
	Board35x45 = &eda.Component{
		CutsOuter: true,

		Cuts: path.Paths{
			path.RoundRect(35, 45, 2.5),
		},

		Nested: eda.Components{
			AlignHole.Clone(2, 30, 0).Clone(2, 0, 40),
		},
	}
)
