// Copyright © 2025 Alex Temnok. All rights reserved.

package boards

import (
	"temnok/pcbc/eda"
	"temnok/pcbc/path"
)

/*

3x2: 22x20
2x2: 35x20
3x1: 22x45
2x1: 35x45
1x1: 75x45

*/

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
