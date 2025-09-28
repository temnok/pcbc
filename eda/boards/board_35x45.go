// Copyright © 2025 Alex Temnok. All rights reserved.

package boards

import (
	"temnok/pcbc/eda"
	"temnok/pcbc/path"
)

/*

1x1: 44x76
1x2: 44x36
1x3: 44x22
2x2: 20x36
2x3: 20x22
2x4: 20x16
3x3: 12x22
3x4: 12x16
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
