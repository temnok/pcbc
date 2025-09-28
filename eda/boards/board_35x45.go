// Copyright © 2025 Alex Temnok. All rights reserved.

package boards

import (
	"temnok/pcbc/eda"
	"temnok/pcbc/path"
)

/*

1x1: 76x44
1x2: 36x44
1x3: 22x44
2x2: 36x20
2x3: 22x20
2x4: 16x20
3x3: 22x12
3x4: 16x12
3x5: 12x12

*/

var (
	Board35x45 = &eda.Component{
		CutsOuter: true,

		Cuts: path.Paths{
			path.RoundRect(35, 45, 2.5),
		},

		Nested: eda.Components{
			AlignHole.CloneX(2, 30).CloneY(2, 40),
		},
	}
)
